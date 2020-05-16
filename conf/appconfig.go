package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var appConfig Config

type Config struct {
	ADMIN_NAME      string `yaml:"admin"`
	MONGO_URI       string `yaml:"uri"`
	ADMIN_PASSWORD  string `yaml:"adminpwd"`
	GIN_MODE        string `yaml:"mode"`
	MONGO_DATABASE  string `yaml:"database"`
	OSS_AK          string `yaml:"ak"`
	OSS_SK          string `yaml:"sk"`
	OSS_BUCKET      string `yaml:"bucket"`
	ALIYUN_ACCOUNT  string `yaml:"aliaccount"`
	AccessKeyID     string `yaml:"akid"`
	AccessKeySecret string `yaml:"aks"`
	Ali_Bucket      string `yaml:"alibucket"`
	Ali_EndPoint    string `yaml:"aliendpoint"`
}

func Load(){
	config, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(config, &appConfig)
}
