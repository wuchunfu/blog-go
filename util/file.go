package util

import (
	"bufio"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"log"
	"mime/multipart"
	"myblog/config"
	"net/http"
	"net/url"
	"os"
)

var (
	FileUtil fileUtils
	tcConf   = config.TcConf
)

type fileUtils struct {
}

func (*fileUtils) UploadFile(file *multipart.FileHeader, path string) string {
	u, _ := url.Parse(tcConf.Url)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  tcConf.SecretID,
			SecretKey: tcConf.SecretKey,
		},
	})
	open, err := file.Open()
	if err != nil {
		log.Println("file open error: ", err)
		return ""
	}
	defer open.Close()
	ctx := context.Background()
	if _, err = client.Object.Put(ctx, path+file.Filename, open, nil); err != nil {
		panic(err)
	}
	return client.Object.GetObjectURL(path + file.Filename).String()
}

func (*fileUtils) WriteFile(fileName string, filePath string, fileContent string) {
	file, err := os.OpenFile(filePath+fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println("file open error: ", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	count, err := writer.WriteString(fileContent)
	writer.Flush()
	if err != nil {
		log.Println("写入错误")
	} else {
		log.Println("写入成功", count)
	}
}

func (*fileUtils) ReadFile(file *multipart.FileHeader) string {
	open, err := file.Open()
	if err != nil {
		log.Println("file open error: ", err)
		return ""
	}
	defer open.Close()
	all, err := io.ReadAll(open)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(all)
}
