package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client interface {
	ListFiles(ctx context.Context, s3Directory string) ([]string, error)
}

type liveClient struct {
	client *s3.Client
}

func (lc *liveClient) ListFiles(ctx context.Context, s3Directory string) ([]string, error) {

	return nil, nil
}
