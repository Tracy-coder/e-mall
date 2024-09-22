// Code generated by hertz generator.

package user

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/biz/handler/middleware"
	logic "github.com/Tracy-coder/e-mall/biz/logic"
	"github.com/Tracy-coder/e-mall/biz/model/user"
	pb "github.com/Tracy-coder/e-mall/biz/model/user"
	"github.com/Tracy-coder/e-mall/configs"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
)

// Register .
// @router /api/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(user.BaseResp)
	var req pb.UserRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = pb.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	if configs.Data().IsProd {
		valid := logic.CaptchaStore.Verify(req.CaptchaID, req.Captcha, true)
		if !valid {
			resp.ErrCode = pb.ErrCode_CaptchaMismatchError
			c.JSON(consts.StatusBadRequest, resp)
			return
		}
	}

	var userRegisterReq domain.UserRegisterReq
	_ = copier.Copy(&userRegisterReq, &req)
	fmt.Println(req)
	err = logic.NewUser(data.Default()).Register(ctx, userRegisterReq)
	if err != nil {
		resp.ErrCode = pb.ErrCode_CreateUserError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = pb.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// Captcha .
// @router /api/captcha [GET]
func Captcha(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.Empty
	resp := new(user.CaptchaInfoResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = pb.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	id, b64s, err := logic.NewCaptcha().GetCaptcha()
	if err != nil {
		resp.ErrCode = pb.ErrCode_CaptchaError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = pb.ErrCode_Success
	resp.ErrMsg = "success"
	resp.CaptchaID = id
	resp.ImgPath = b64s

	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router /api/user/info [POST]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.Empty
	resp := new(user.UserInfoResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = pb.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	v, exist := c.Get("userID")
	if !exist || v == nil {
		c.JSON(consts.StatusUnauthorized, "Unauthorized")
		return
	}
	i, err := strconv.Atoi(v.(string))
	if err != nil {
		c.JSON(consts.StatusUnauthorized, "Unauthorized,"+err.Error())
		return
	}
	userID := uint64(i)
	user, err := logic.NewUser(data.Default()).UserInfo(ctx, userID)
	if err != nil {
		resp.ErrCode = pb.ErrCode_GetUserInfoError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	resp.ID = user.ID
	resp.Username = user.Username
	resp.Status = uint64(user.Status)
	resp.Email = user.Email
	resp.Mobile = user.Mobile
	resp.Avatar = user.Avatar
	resp.Nickname = user.Nickname
	resp.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")

	resp.ErrCode = pb.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)

}

// GTLogin .
// @router /api/GT/login [GET]
func GTLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.Empty
	resp := new(pb.GTLoginResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = pb.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	url := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=http://localhost:8888/api/v1/github/login/callback", os.Getenv("GITHUB_KEY"))
	resp.ErrCode = pb.ErrCode_Success
	resp.LoginURL = url
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// GTLoginCallback .
// @router /api/github/login/callback [POST]
func GTLoginCallback(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GTLoginCallbackReq
	resp := new(pb.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = pb.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	userInfo, err := logic.NewUser(data.Default()).GTLoginCallback(ctx, req.Code)
	if err != nil {
		fmt.Println(err)
	}

	c.Set("userInfo", *userInfo)
	fmt.Println(userInfo)

	ctx = context.WithValue(ctx, "OAuth", 1) //TODO
	middleware.GetJWTMiddleware(configs.Data(), data.Default()).LoginHandler(ctx, c)
}

// ValidEmail .
// @router /api/user/valid-email [POST]
func ValidEmail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ValidEmailReq
	resp := new(user.ValidEmailResp)
	err = c.BindAndValidate(&req)
	fmt.Println(req)
	if err != nil {
		resp.ErrCode = user.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		return
	}

	userID, err := strconv.ParseUint(c.Value("userID").(string), 10, 64)
	if err != nil {
		resp.ErrCode = user.ErrCode_GetUserIdError
		resp.ErrMsg = "failed to get userID from context"
		return
	}
	if req.Op == user.EmailOperation_Binding {
		fmt.Println("bind")
		err := logic.NewUser(data.Default()).BindEmail(ctx, userID, req.Email)
		if err != nil {
			fmt.Println(err.Error())
			resp.ErrCode = user.ErrCode_BindEmailError
			resp.ErrMsg = err.Error()
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
	} else {
		fmt.Println("unbind")
		err := logic.NewUser(data.Default()).UnbindEmail(ctx, userID)
		if err != nil {
			resp.ErrCode = user.ErrCode_UnbindEmailError
			resp.ErrMsg = err.Error()
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
	}
	resp.ErrCode = user.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// VerifyEmail .
// @router /api/v1/verify_email [GET]
func VerifyEmail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.VerifyEmailReq
	resp := new(user.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = pb.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	fmt.Println(req)
	err = logic.NewUser(data.Default()).VerfifyEmail(ctx, req.EmailId, req.SecretCode)
	if err != nil {
		resp.ErrCode = pb.ErrCode_VerifyEmailError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	resp.ErrCode = pb.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}
