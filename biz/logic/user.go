package logic

import (
	"context"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/pkg/encrypt"
	"github.com/pkg/errors"
)

type User struct {
	Data *data.Data
}

func NewUser(data *data.Data) domain.User {
	return &User{
		Data: data,
	}
}

func (u *User) Register(ctx context.Context, req domain.UserRegisterReq) error {
	password, _ := encrypt.BcryptEncrypt(req.Password)
	_, err := u.Data.DBClient.User.Create().
		SetEmail(req.Email).
		SetUsername(req.Username).
		SetPassword(password).
		SetNickname(req.Nickname).
		Save(ctx)
	if err != nil {
		err = errors.Wrap(err, "create user failed")
		return err
	}

	return nil
}
