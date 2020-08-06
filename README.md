## cloudstroe

Cloudtroe is an integrated object storage package that implements obs (Huawei cloud storage), cos (Tencent cloud storage), oss (Ali cloud storage), and ministore related interfaces. About ministore you can visit the official [websiteMinio](https://docs.min.io/cn)

cloudstroe 是一个集成的对象存储包，实现了obs(华为云存储),cos(腾讯云存储),oss(阿里云存储),ministore相关接口。关于ministore你可以访问官网

## Getting started

```bash
go get github.com/fromiuan/cloudstroe
```

You're good to go, happy coding! 😁

## Usage

```bash
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/fromiuan/cloudstroe"
)

func main(){
	cfg := cloudstroe.CosConfig{
		AccessKey: "accesskey_id",
		SecretKey: "secret_access_key",
		Bucket:    "bucket_name",
		Endpoint:  "endpoint",
		Domain:    "domain",
	}
	client, err = cloudstroe.New(cloudstroe.Config{Driver: "cos", Config: cfg})
	if err != nil {
		panic("cloud init error")
	}

	// upload file
	images := "./images.png"
	fileInfo,err := os.Stat(images)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Open("./images.png")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	err := client.UploadObject(objectName, file, fileInfo.Size())
	if err != nil {
		fmt.Println(err)
	}

	// get file
	object, err := getMiniClient().GetObject(myobject)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer object.Close()

	file, err := os.Open("./images.png")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	n,err := io.Copy(file,object)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
```