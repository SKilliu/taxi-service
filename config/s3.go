package config

import (
	"github.com/caarlos0/env"
)

type S3 struct {
	AccessKey string `env:"TAXI_SERVICE_S3_ACCESS_KEY"`
	SecretKey string `env:"TAXI_SERVICE_S3_SECRET_KEY"`
	Url       string `env:"TAXI_SERVICE_S3_URL"`
	Bucket    string `env:"TAXI_SERVICE_S3_BUCKET"`
}

func (config *ConfigImpl) S3() *s3.Client {
	if config.s3 != nil {
		return config.s3
	}

	config.Lock()
	defer config.Unlock()

	s3Config := new(S3)
	if err := env.Parse(s3Config); err != nil {
		config.log.WithError(err).Error("failed to parse minio settings")
		panic(err)
	}

	client, err := s3.New(
		config.log,
		s3Config.AccessKey,
		s3Config.SecretKey,
		s3Config.Bucket,
		s3Config.Url,
	)
	if err != nil {
		config.log.WithError(err).Error("failed to create default minio client")
		panic(err)
	}

	config.s3 = client
	return config.s3
}
