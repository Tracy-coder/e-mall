package logic

import (
	"context"
	"log"
	"net/url"
	"time"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type OSS struct {
	minio *minio.Client
}

const (
	accessKey    = "FEmK32D4DvwG7Q1WJ465"
	accessSecret = "6ECeIYCEGWZgCcrUw8ZHjo6qPfQnK1BiqpbLOIVZ"
	endPoint     = "127.0.0.1:9000" // minio地址,不能加http
	useSSL       = false            // 是否使用https进行通信
)

var (
	minioClient *minio.Client
)

func init() {
	var err error
	minioClient, err = minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, accessSecret, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal("minio client create fail, err %+v", err)
	}
}

func NewOSS() domain.OSS {
	return &OSS{
		minio: minioClient,
	}
}

func (o *OSS) UploadAvatar(ctx context.Context, bucketName string, objectName string) (*domain.UploadAvatarResp, error) {
	resp := new(domain.UploadAvatarResp)
	u, err := minioClient.PresignedPutObject(ctx, bucketName, objectName, time.Hour)
	if err != nil {
		return nil, err
	}
	resp.Put = u.Host + u.Path + "?" + u.RawQuery
	u, err = minioClient.PresignedGetObject(ctx, bucketName, objectName, time.Hour, url.Values{})
	if err != nil {
		return nil, err
	}
	resp.Get = u.Host + u.Path + "?" + u.RawQuery
	return resp, nil
}
