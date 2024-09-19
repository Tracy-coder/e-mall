// from: https://github.com/chenghonour/formulago

package middleware

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/biz/logic"
	"github.com/Tracy-coder/e-mall/configs"
	Data "github.com/Tracy-coder/e-mall/data"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cockroachdb/errors"
	"github.com/hertz-contrib/jwt"
)

type jwtLogin struct {
	Username  string `form:"username,required" json:"username,required"`   //lint:ignore SA5008 ignoreCheck
	Password  string `form:"password,required" json:"password,required"`   //lint:ignore SA5008 ignoreCheck
	Captcha   string `form:"captcha,required" json:"captcha,required"`     //lint:ignore SA5008 ignoreCheck
	CaptchaID string `form:"captchaID,required" json:"captchaId,required"` //lint:ignore SA5008 ignoreCheck
}

// jwt identityKey
var (
	identityKey   = "jwt-id"
	jwtMiddleware = new(jwt.HertzJWTMiddleware)
)

// GetJWTMiddleware returns a new JWT middleware.
func GetJWTMiddleware(c configs.Config, d *Data.Data) *jwt.HertzJWTMiddleware {
	jwtMiddleware, err := newJWT(c, d)
	if err != nil {
		hlog.Fatal(err, "JWT Init Error")
	}
	return jwtMiddleware
}

func newJWT(config configs.Config, db *Data.Data) (jwtMiddleware *jwt.HertzJWTMiddleware, err error) {
	// the jwt middleware
	jwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "e-mall",
		Key:         []byte("test"),
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// take map which have roleID, userID as Payload
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					identityKey: v, //其实就是一股脑把这个map丢进来
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			payloadMap, ok := claims[identityKey].(map[string]interface{})
			if !ok {
				hlog.Error("get payloadMap error", "claims data:", claims[identityKey])
				return nil
			}
			c.Set("userID", payloadMap["userID"])
			// ID是为了给Authorizator用
			return payloadMap
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			oauthLogin := ctx.Value("OAuth") == 1
			if !oauthLogin {
				res := new(domain.UserLoginResp)
				// normal jwtLogin
				var loginVal jwtLogin
				if err := c.BindAndValidate(&loginVal); err != nil {
					return "", err
				}
				// verify captcha while IsProd is true
				// 开发模式不做检查
				if config.IsProd {
					valid := logic.CaptchaStore.Verify(loginVal.CaptchaID, loginVal.Captcha, true)
					if !valid {
						return nil, errors.New("invalid captcha")
					}
				}
				// Login
				username := loginVal.Username
				password := loginVal.Password
				res, err = logic.NewUser(db).Login(ctx, username, password)
				if err != nil {
					hlog.Error(err, "jwtLogin error")
					return nil, err
				}
				// return the payload
				// take str roleID, userID into PayloadMap
				payloadMap := make(map[string]interface{})
				payloadMap["userID"] = strconv.Itoa(int(res.UserID))
				return payloadMap, nil
			} else {
				fmt.Println("OAuth login")
				userInfo, ok := c.Value("userInfo").(domain.OauthUserInfo)
				fmt.Println(userInfo)
				if !ok {
					return nil, errors.New("get ctx value failed")
				}
				res, err := logic.NewUser(db).OAuthLogin(ctx, userInfo.ID)
				if err != nil {
					return nil, err
				}
				payloadMap := make(map[string]interface{})
				payloadMap["userID"] = strconv.Itoa(int(res.UserID))
				return payloadMap, nil
			}

		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
	})

	return
}
