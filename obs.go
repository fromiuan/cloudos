package cloudos

import (
	"errors"
	"io"
)

var obsDriver = &OBS{}

type OBS struct {
}

func (o *OBS) Init(cfg interface{}) error {
	return errors.New("params must MiniConfig struct")
}

func (o *OBS) UploadObject(name string, reader io.Reader, size int64) error {
	return nil
}

func (o *OBS) GetObject(name string) (io.ReadCloser, error) {
	return nil, nil
}

func init() {
	Register("obs", obsDriver)
}
