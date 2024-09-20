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
	"github.com/Tracy-coder/e-mall/data/ent/email"
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

func (u *User) GTLoginCallback(ctx context.Context, code string) (*domain.OauthUserInfo, error) {
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
		return nil, err
	}
	fmt.Println(token)
	var response *http.Response
	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var userInfo *domain.OauthUserInfo
	err = json.Unmarshal(contents, &userInfo)
	if err != nil {
		return nil, err
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
			return nil, err
		}
	}
	fmt.Println(userDB)
	return userInfo, nil
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

func (u *User) OAuthLogin(ctx context.Context, githubID uint64) (*domain.UserLoginResp, error) {
	userEnt, err := u.Data.DBClient.User.Query().Where(user.GithubID(githubID)).First(ctx)
	fmt.Println(userEnt)
	if err != nil {
		err = errors.Wrap(err, "get user failed")
		return nil, err
	}
	return &domain.UserLoginResp{
		Username: userEnt.Username,
		UserID:   userEnt.ID,
	}, nil
}
func (u *User) BindEmail(ctx context.Context, id uint64, email_address string) error {
	fmt.Println(id, email_address)
	res, err := u.Data.DBClient.Email.Query().Where(email.Email(email_address), email.UserID(id), email.IsVerified(true)).First(ctx)
	if res != nil {
		return errors.New("There is already a verified email address, please unbind it first")
	}
	err = u.Data.DBClient.User.Update().SetEmail(email_address).Where(user.ID(id)).Exec(ctx)
	if err != nil {
		return err
	}
	secretCode := generateSecret()
	verifyEmail, err := u.Data.DBClient.Email.Create().
		SetEmail(email_address).
		SetIsVerified(false).
		SetUserID(id).
		SetSecret(secretCode).
		Save(ctx)
	subject := "Welcome to wzw's E-mall"
	// TODO: replace this URL with an environment variable that points to a front-end page
	verifyUrl := fmt.Sprintf("http://localhost:8888/api/v1/verify_email?email_id=%d&secret_code=%s",
		verifyEmail.ID, verifyEmail.Secret)
	content := fmt.Sprintf(`Hello!<br/>
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, verifyUrl)
	to := []string{email_address}

	err = NewGmailSender("e-mall", os.Getenv("EMAIL_ADDR"), os.Getenv("EMAIL_KEY")).SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}
	// TODO:kafka
	return nil
}

func (u *User) UnbindEmail(ctx context.Context, id uint64) error {
	//TODO:transactions
	_, err := u.Data.DBClient.Email.Delete().Where(email.UserID(id)).Exec(ctx)

	if err != nil {
		return err
	}
	return u.Data.DBClient.User.Update().ClearEmail().Where(user.ID(id)).Exec(ctx)
}

func generateSecret() string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 32)
	rand.Seed(uint64(time.Now().UnixNano()))
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func (u *User) VerfifyEmail(ctx context.Context, emailID uint64, secretCode string) error {
	emailDB, err := u.Data.DBClient.Email.Query().Where(email.ID(emailID)).First(ctx)
	if err != nil {
		return err
	}
	if emailDB.Secret != secretCode {
		return errors.New("secret code mismatch")
	}
	return u.Data.DBClient.Email.Update().Where(email.ID(emailID)).SetIsVerified(true).Exec(ctx)
}
