package cloudos

import (
	"fmt"
	"io"
)

type Config struct {
	Driver string // 存储名称(cos,oss,obs,minio)
	Config interface{}
}

type Driver interface {
	Init(cfg interface{}) error
	UploadObject(string, io.Reader, int64) error
	GetObject(string) (io.ReadCloser, error)
}

var driverGlobal = make(map[string]Driver)

func New(cfg Config) (Driver, error) {
	driver, b := driverGlobal[cfg.Driver]
	if b {
		return nil, fmt.Errorf("driver[%s] not support", cfg.Driver)
	}
	if err := driver.Init(cfg.Config); err != nil {
		return nil, err
	}
	return driver, nil
}

func Register(key string, d Driver) {
	if d == nil {
		panic("Register driver is error")
	}
	if _, ok := driverGlobal[key]; ok {
		panic("Register called twice for driver" + key)
	}
	driverGlobal[key] = d
}
