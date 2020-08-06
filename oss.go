package cloudos

import (
	"errors"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossDriver = &OSS{}

type OSS struct {
	Client *oss.Client
	Cfg    OssConfig
}

func (o *OSS) Init(cfg interface{}) error {
	if config, ok := cfg.(OssConfig); ok {
		client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
		if err != nil {
			return err
		}
		o.Client = client
		o.Cfg = config
		return nil
	}
	return errors.New("params must MiniConfig struct")
}

func (o *OSS) UploadObject(name string, reader io.Reader, size int64) error {
	bucket, err := o.Client.Bucket(o.Cfg.Bucket)
	if err != nil {
		return err
	}
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	err = bucket.PutObject(name, reader, storageType, objectAcl)
	if err != nil {
		return err
	}
	return nil
}

func (o *OSS) GetObject(name string) (io.ReadCloser, error) {
	bucket, err := o.Client.Bucket(o.Cfg.Bucket)
	if err != nil {
		return err
	}

	body, err := bucket.GetObject(name)
	if err != nil {
		return err
	}
	return body, nil
}

func init() {
	Register("oss", ossDriver)
}
