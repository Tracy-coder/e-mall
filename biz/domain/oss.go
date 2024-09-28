package domain

import "context"

type OSS interface {
	UploadAvatar(ctx context.Context, bucketName string, objectName string) (*UploadAvatarResp, error)
}

type UploadAvatarResp struct {
	Put string
	Get string
}
