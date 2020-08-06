package cloudos

import (
	"errors"
	"io"

	"github.com/fromiuan/cloudstroe/lib/obs"
)

var obsDriver = &OBS{}

type OBS struct {
	Client *obs.ObsClient
	Cfg    ObsConfig
}

func (o *OBS) Init(cfg interface{}) error {
	if config, ok := cfg.(ObsConfig); ok {
		client, err := obs.New(config.AccessKey, config.SecretKey, config.Endpoint)
		if err != nil {
			return err
		}
		o.Client = client
		o.Cfg = config
		return nil
	}
	return errors.New("params must CosConfig struct")
}

func (o *OBS) UploadObject(name string, reader io.Reader, size int64) error {
	input := &obs.PutObjectInput{}
	input.Body = reader
	input.Bucket = o.Cfg.Bucket
	input.Key = name
	_, err := o.Client.PutObject(input)
	if err != nil {
		return err
	}
	return nil
}

func (o *OBS) GetObject(name string) (io.ReadCloser, error) {
	input := &obs.GetObjectInput{}
	input.Bucket = o.Cfg.Bucket
	input.Key = name
	output, err := o.Client.GetObject(input)
	if err != nil {
		return nil, err
	}
	return output.Body, nil
}

func init() {
	Register("obs", obsDriver)
}
