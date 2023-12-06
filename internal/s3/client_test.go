package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsCredentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/docker/go-connections/nat"
	"github.com/minio/minio-go/v7"
	minioCredentials "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"path"
	"testing"
	"time"
)

const (
	bucketName   = "bucket"
	bucketPrefix = "path/to/files"
	region       = "us-east-1"
)

type s3ConnectionProps struct {
	endpoint        string
	accessKeyID     string
	secretAccessKey string
}

/*
resolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
    return aws.Endpoint{
        PartitionID:       "aws",
        URL:               "http://127.0.0.1:9000",
        SigningRegion:     "us-east-2",
        HostnameImmutable: true,
    }, nil
})

conn := s3.NewFromConfig(aws.Config{
    Region:           "us-east-2",
    Credentials:      credentials.NewStaticCredentialsProvider("minioadmin", "minioadmin", ""),
    EndpointResolver: resolver,
}, func(o *s3.Options) {
    o.UsePathStyle = true
})
*/

func Test_ListFiles_ReturnsListOfFilesFromTheBucket(t *testing.T) {
	s3Files := map[string]string{
		"/file1.txt": "foo",
		"/file2.txt": "bar",
		"/file3.txt": "baz",
	}

	ctx := context.Background()
	connectionProps := setupS3Container(t, ctx, s3Files)

	config := aws.Config{
		Region:       region,
		Credentials:  awsCredentials.NewStaticCredentialsProvider(connectionProps.accessKeyID, connectionProps.secretAccessKey, ""),
		BaseEndpoint: aws.String(connectionProps.endpoint),
	}

	client := NewClient(config, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	files, err := client.ListFiles(ctx, "s3://"+path.Join(bucketName, bucketPrefix))
	require.NoError(t, err)
	require.NotEmpty(t, files)
}

func setupS3Container(t *testing.T, ctx context.Context, files map[string]string) s3ConnectionProps {
	accessKeyId := "root"
	secretAccessKey := "password"
	minioContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "minio/minio:latest",
			ExposedPorts: []string{"9000", "9090"},
			WaitingFor:   wait.ForLog("MinIO Object Storage Server").WithStartupTimeout(10 * time.Second),
			Env:          map[string]string{"MINIO_ROOT_USER": accessKeyId, "MINIO_ROOT_PASSWORD": secretAccessKey},
			Cmd:          []string{"server", "/data"},
		},
		Started: true,
	})
	if err != nil {
		t.Error(err)
	}
	t.Cleanup(func() {
		if err := minioContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	})

	host, err := minioContainer.Host(ctx)
	require.NoError(t, err)
	minioPort, err := nat.NewPort("", "9000")
	require.NoError(t, err)
	port, err := minioContainer.MappedPort(ctx, minioPort)
	require.NoError(t, err)
	endpoint := fmt.Sprintf("%s:%s", host, port.Port())

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  minioCredentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: false,
	})
	require.NoError(t, err)

	// Create the bucket
	if err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if !(errBucketExists == nil && exists) {
			log.Fatalf("failed to create bucket %q: %v", bucketName, err)
		}
	}

	// Put files into the created bucket
	for fileName, fileContent := range files {
		putObjectOpts := minio.PutObjectOptions{ContentType: "application/octet-stream"}
		reader := bytes.NewReader([]byte(fileContent))
		_, err = minioClient.PutObject(ctx, bucketName, bucketPrefix+fileName, reader, reader.Size(), putObjectOpts)
		require.NoError(t, err)
	}

	return s3ConnectionProps{
		endpoint:        "http://" + endpoint,
		accessKeyID:     accessKeyId,
		secretAccessKey: secretAccessKey,
	}
}
