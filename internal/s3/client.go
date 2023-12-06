package s3

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"strings"
)

const (
	s3Prefix = "s3://"
)

type Client interface {
	ListFiles(ctx context.Context, s3Directory string) ([]string, error)
}

func NewClientWithDefaultConfig() Client {
	return NewClient(aws.Config{Region: "us-east-1"})
}

func NewClient(sdkConfig aws.Config, opts ...func(*s3.Options)) Client {
	return &liveClient{client: s3.NewFromConfig(sdkConfig, opts...)}
}

type liveClient struct {
	client *s3.Client
}

func (lc *liveClient) ListFiles(ctx context.Context, s3Directory string) ([]string, error) {
	listObjectsInput, err := deriveListObjectsInput(s3Directory)
	if err != nil {
		return nil, err
	}

	files := make([]string, 0)
	res, err := lc.client.ListObjectsV2(ctx, listObjectsInput)
	if err != nil {
		return nil, err
	}

	for _, obj := range res.Contents {
		files = append(files, *obj.Key)
	}
	return files, nil
}

func deriveListObjectsInput(scoresDistributionLocation string) (*s3.ListObjectsV2Input, error) {
	if !strings.Contains(scoresDistributionLocation, s3Prefix) {
		return nil, errors.New("invalid s3 path for scores distribution: " + scoresDistributionLocation)
	}

	s3Path := strings.ReplaceAll(scoresDistributionLocation, s3Prefix, "")
	bucketNameAndPrefix := strings.SplitN(s3Path, "/", 2)

	if len(bucketNameAndPrefix) != 2 {
		return nil, errors.New("invalid s3 path for scores distribution: " + scoresDistributionLocation)
	}

	return &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketNameAndPrefix[0]),
		Prefix: aws.String(bucketNameAndPrefix[1]),
	}, nil
}
