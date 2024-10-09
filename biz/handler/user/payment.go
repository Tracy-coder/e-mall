// Code generated by hertz generator.

package user

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Tracy-coder/e-mall/biz/logic"
	base "github.com/Tracy-coder/e-mall/biz/model/base"
	user "github.com/Tracy-coder/e-mall/biz/model/user"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// InitPay .
// @router /api/v1/user/payments [GET]
func InitPay(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.InitPayReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(base.BaseResp)
	userID, err := getUserID(c)
	if err != nil {
		resp.ErrCode = base.ErrCode_UnauthorizedError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusUnauthorized, resp)
		return
	}
	url, err := logic.NewPayment(data.Default()).Init(ctx, req.OrderNum, userID)
	if err != nil {
		resp.ErrMsg = err.Error()
		resp.ErrCode = base.ErrCode_InitPayError
	}
	c.Redirect(http.StatusTemporaryRedirect, []byte(url.String()))
}

// NotifyPay .
// @router /api/v1/alipay/notify [POST]
func NotifyPay(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.Empty
	err = c.BindAndValidate(&req)
	resp := new(base.BaseResp)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	v := url.Values{}
	c.Request.PostArgs().VisitAll(func(key, value []byte) { v.Add(string(key), string(value)) })
	err = logic.NewPayment(data.Default()).Notify(ctx, v)
	if err != nil {
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// CallbackPay .
// @router /api/v1/alipay/callback [GET]
// TODO:callback函数未做任何处理
func CallbackPay(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// a := url.Values{}
	// c.Request.PostArgs().VisitAll(func(key, value []byte) { a.Add(string(key), string(value)) })
	// fmt.Println(a)
	resp := new(base.BaseResp)
	resp.ErrMsg = "可以关闭本页面~"
	c.JSON(consts.StatusOK, resp)
}
