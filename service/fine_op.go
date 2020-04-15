package service

import (
	"log"
	"makespace-remaster/conf"
	"os"
)

func FileUpload(filepath string) {
	bucket, err := conf.OssClient.Bucket(conf.Bucket)
	if err != nil {
		log.Fatal(err)
	}
	fd,err := os.Open(filepath)
	if err != nil{
		log.Fatal(err)
	}
	defer fd.Close()
	err = bucket.PutObject("main.txt",fd)
	if err != nil{
		log.Fatal(err)
	}
}
