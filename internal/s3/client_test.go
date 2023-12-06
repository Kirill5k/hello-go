package s3

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
