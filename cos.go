package cloudos

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type COS struct {
	Client *cos.Client
}

var cosDriver = &COS{}

func (c *COS) Init(cfg interface{}) error {
	if config, ok := cfg.(CosConfig); ok {
		u, err := url.Parse(config.Region)
		if err != nil {
			return err
		}
		client := cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  config.SecretID,
				SecretKey: config.SecretKey,
			},
		})
		if client == nil {
			return errors.New("client is nil")
		}
		c.Client = client
		return nil
	}
	return errors.New("params must CosConfig struct")

}

func (c *COS) UploadObject(name string, reader io.Reader, size int64) error {
	_, err := c.Client.Object.Put(context.Background(), name, reader, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *COS) GetObject(name string) (io.ReadCloser, error) {
	object, err := c.Client.Object.Get(context.Background(), name, nil)
	if err != nil {
		return nil, err
	}
	if object.StatusCode != http.StatusOK {
		return nil, errors.New("error staus")
	}
	return object.Body, nil
}

func init() {
	Register("cos", cosDriver)
}
