package logic

import (
	"context"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/data/ent/user"
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

func (u *User) Login(ctx context.Context, username string, password string) (res *domain.UserLoginResp, err error) {
	user, err := u.Data.DBClient.User.Query().Where(user.Username(username), user.Status(1)).Only(ctx)
	if err != nil {
		return nil, err
	}
	if ok := encrypt.BcryptCheck(password, user.Password); !ok {
		err = errors.New("wrong password")
		return nil, err
	}
	res = new(domain.UserLoginResp)
	res.Username = username
	res.UserID = user.ID

	return
}
