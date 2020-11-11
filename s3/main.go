package s3

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

type Client struct {
	minio  *minio.Client
	Log    *logrus.Entry
	bucket string
}

func New(log *logrus.Entry, accessKey, secretKey, bucket, url string) (*Client, error) {
	client, err := minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		minio:  client,
		Log:    log,
		bucket: bucket,
	}, nil
}

func (c Client) URL(fileName string) string {
	return fmt.Sprintf("http://localhost:9000/%s/%s", c.bucket, fileName)
}
