package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/data/ent/user"
	"github.com/Tracy-coder/e-mall/pkg/encrypt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"golang.org/x/exp/rand"
	"golang.org/x/oauth2"
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

func (u *User) UserInfo(ctx context.Context, id uint64) (userInfo domain.UserInfo, err error) {
	userEnt, err := u.Data.DBClient.User.Query().Where(user.IDEQ(id)).First(ctx)
	if err != nil {
		err = errors.Wrap(err, "get user failed")
		return userInfo, err
	}
	// copy to UserInfo struct
	err = copier.Copy(&userInfo, &userEnt)
	return
}

func (u *User) GTLoginCallback(ctx context.Context, code string) error {
	config := oauth2.Config{
		ClientID:     os.Getenv("GITHUB_KEY"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		RedirectURL: "http://localhost:8888/api/github/login/callback",
	}
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return err
	}
	fmt.Println(token)
	var response *http.Response
	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var userInfo *domain.OauthUserInfo
	err = json.Unmarshal(contents, &userInfo)
	if err != nil {
		return err
	}
	userDB, err := u.Data.DBClient.User.Query().Where(user.GithubID(userInfo.ID), user.Status(1)).First(ctx)
	if err != nil {
		newUsername := userInfo.Username
		_, errNoUser := u.Data.DBClient.User.Query().Where(user.Username(newUsername), user.Status(1)).First(ctx)
		for errNoUser == nil {
			newUsername = generateRandomName(userInfo.Username)
			_, errNoUser = u.Data.DBClient.User.Query().Where(user.Username(newUsername), user.Status(1)).First(ctx)
		}
		userDB, err = u.Data.DBClient.User.Create().
			SetUsername(newUsername).
			SetNickname(userInfo.Nickname).
			SetGithubID(userInfo.ID).
			Save(ctx)
		if err != nil {
			return err
		}
	}
	fmt.Println(userDB)
	return nil
}

func generateRandomName(name string) (newName string) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 6)
	rand.Seed(uint64(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return name + "_" + string(b)
}
