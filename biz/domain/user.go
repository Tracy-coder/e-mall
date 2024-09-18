package domain

import (
	"context"
	"time"
)

type UserRegisterReq struct {
	ID       uint64
	Email    string
	Username string
	Password string
	Nickname string
}

type User interface {
	Register(ctx context.Context, req UserRegisterReq) error
	UserInfo(ctx context.Context, id uint64) (userInfo UserInfo, err error)
	Login(ctx context.Context, username string, password string) (res *UserLoginResp, err error)
	GTLoginCallback(ctx context.Context, code string) error
	// Update(ctx context.Context, req CreateOrUpdateUserReq) error
	// ChangePassword(ctx context.Context, userID uint64, oldPassword, newPassword string) error
	// UserInfo(ctx context.Context, id uint64) (userInfo *UserInfo, err error)
	// List(ctx context.Context, req UserListReq) (userList []*UserInfo, total int, err error)
	// UpdateUserStatus(ctx context.Context, id uint64, status uint64) error
	// DeleteUser(ctx context.Context, id uint64) error
	// UpdateProfile(ctx context.Context, req UpdateUserProfileReq) error
}

type UserLoginResp struct {
	UserID   uint64
	Username string
}

type UserInfo struct {
	ID        uint64
	Status    uint8
	Username  string
	Password  string
	Nickname  string
	Mobile    string
	Email     string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OauthUserInfo struct {
	Username string `json:"login"`
	Avatar   string `json:"avatar_url"`
	Nickname string `json:"name"`
	ID       uint64 `json:"id"`
}
