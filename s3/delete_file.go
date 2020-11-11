package s3

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (c Client) DropFile(filename string) error {
	return c.minio.RemoveObject(context.Background(), c.bucket, filename, minio.RemoveObjectOptions{})
}
