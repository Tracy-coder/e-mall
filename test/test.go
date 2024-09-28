package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	accessKey := "FEmK32D4DvwG7Q1WJ465"
	accessSecret := "6ECeIYCEGWZgCcrUw8ZHjo6qPfQnK1BiqpbLOIVZ"
	endPoint := "127.0.0.1:9000" // minio地址,不能加http
	useSSL := false              // 是否使用https进行通信

	// 初始化minio客户端
	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, accessSecret, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal("minio client create fail, err %+v", err)
	}
	filePath := "/home/yjy/projects/cmall/e-mall/test/1.jpg"
	uploadInfo, err := minioClient.FPutObject(context.Background(), "avatar", "111.jpg", filePath,
		minio.PutObjectOptions{
			ContentType: "application/csv",
		})
	if err != nil {
		fmt.Println("Fput failed")
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)
	minioClient.PresignedPutObject()
}
