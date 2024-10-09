// Code generated by hertz generator.

package user

import (
	"context"
	"fmt"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/biz/logic"
	base "github.com/Tracy-coder/e-mall/biz/model/base"
	user "github.com/Tracy-coder/e-mall/biz/model/user"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
)

// CreateOrder .
// @router /api/v1/user/orders [POST]
func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateOrderReq
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	var createOrderReq domain.CreateOrderReq
	err = copier.Copy(&createOrderReq, &req)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	userID, err := getUserID(c)
	if err != nil {
		resp.ErrCode = base.ErrCode_UnauthorizedError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusUnauthorized, resp)
		return
	}
	createOrderReq.UserID = userID
	err = logic.NewOrder(data.Default()).CreateOrder(ctx, createOrderReq)

	if err != nil {
		resp.ErrCode = base.ErrCode_CreateOrderError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// ListOrder .
// @router /api/v1/user/orders [GET]
func ListOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ListOrderReq
	resp := new(user.ListOrderResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	userID, err := getUserID(c)
	if err != nil {
		resp.ErrCode = base.ErrCode_UnauthorizedError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusUnauthorized, resp)
		return
	}
	var listOrderReq domain.ListOrderReq
	err = copier.Copy(&listOrderReq, &req)
	listOrderReq.UserID = userID
	info, total, err := logic.NewOrder(data.Default()).ListOrder(ctx, listOrderReq)
	if err != nil {
		resp.ErrCode = base.ErrCode_ListOrderError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	orderInfo := make([]*user.OrderInfo, len(info))
	err = copier.Copy(&orderInfo, &info)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	resp.Info = orderInfo
	resp.Total = int32(total)
	c.JSON(consts.StatusOK, resp)
}

// ShowOrder .
// @router /api/v1/user/order [GET]
func ShowOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ShowOrderReq
	resp := new(user.ShowOrderResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	info, err := logic.NewOrder(data.Default()).ShowOrder(ctx, req.OrderNum)
	if err != nil {
		resp.ErrCode = base.ErrCode_ShowOrderError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	orderInfo := new(user.OrderInfo)
	err = copier.Copy(orderInfo, &info)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	resp.Info = orderInfo
	c.JSON(consts.StatusOK, resp)
}

// ShowCount .
// @router /api/v1/user/count [GET]
func ShowCount(ctx context.Context, c *app.RequestContext) {
	var err error

	resp := new(user.ShowCountResp)
	userID, err := getUserID(c)
	if err != nil {
		resp.ErrCode = base.ErrCode_UnauthorizedError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusUnauthorized, resp)
		return
	}
	res, err := logic.NewOrder(data.Default()).ShowCount(ctx, userID)
	if err != nil {
		resp.ErrCode = base.ErrCode_ShowCountError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	fmt.Println(res)
	err = copier.Copy(resp, res)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}
