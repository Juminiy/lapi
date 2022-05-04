package storage

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"lapi/config"
)

var (
	OSSClient *oss.Client
)
func OSSConnect() error {
	client,err := oss.New(config.Config("OSS_ENDPOINT"),config.Config("OSS_ACCESS_KEY_ID"),config.Config("OSS_ACCESS_KEY_SECRET"),oss.UseCname(false))
	if err != nil {
		return err
	}
	OSSClient = client
	return nil
}
