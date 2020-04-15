package conf

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var (
	OssClient *oss.Client
	Bucket string
)

func OssInit() error {
	endPoint := os.Getenv("Ali_EndPoint")
	acessKeyId := os.Getenv("AccessKeyID")
	accessKeySecret := os.Getenv("AccessKeySecret")
	client, err := oss.New(endPoint, acessKeyId, accessKeySecret)
	if err != nil {
		return err
	}
	OssClient = client
	Bucket = os.Getenv("Ali_Bucket")
	return nil
}
