package cloudos

import (
	"errors"
	"io"
)

var ossDriver = &OSS{}

type OSS struct {
}

func (o *OSS) Init(cfg interface{}) error {
	return errors.New("params must MiniConfig struct")
}

func (o *OSS) UploadObject(name string, reader io.Reader, size int64) error {
	return nil
}

func (o *OSS) GetObject(name string) (io.ReadCloser, error) {
	return nil, nil
}

func init() {
	Register("oss", ossDriver)
}
