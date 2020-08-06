package cloudos

import (
	"errors"
	"io"

	"github.com/minio/minio-go/v6"
)

type Minio struct {
	BucketName string
	Client     *minio.Client
}

var minioDriver = &Minio{}

func (m *Minio) Init(cfg interface{}) error {
	if config, ok := cfg.(MinioConfig); ok {
		client, err := minio.New(config.Endpoint, config.AccessKey, config.SecretKey, config.SSL)
		if err != nil {
			return err
		}
		m.Client = client
		m.BucketName = config.Bucket
		return nil
	}
	return errors.New("params must MiniConfig struct")
}

func (m *Minio) UploadObject(name string, reader io.Reader, size int64) error {
	n, err := m.Client.PutObject(m.BucketName, name, reader, size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil || n <= 0 {
		return err
	}
	return nil
}

func (m *Minio) GetObject(name string) (io.ReadCloser, error) {
	object, err := m.Client.GetObject(m.BucketName, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func init() {
	Register("minio", minioDriver)
}
