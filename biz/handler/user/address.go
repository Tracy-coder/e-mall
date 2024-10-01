// Code generated by hertz generator.

package user

import (
	"context"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/biz/logic"
	base "github.com/Tracy-coder/e-mall/biz/model/base"
	user "github.com/Tracy-coder/e-mall/biz/model/user"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
)

// CreateAddress .
// @router /api/v1/user/addresses [POST]
func CreateAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateAddressReq
	err = c.BindAndValidate(&req)
	resp := new(user.CreateAddressResp)
	if err != nil {
		resp.ErrCode = base.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	createAddressReq := new(domain.CreateAddressReq)
	err = copier.Copy(createAddressReq, &req)
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
	createAddressReq.UserID = userID
	info, err := logic.NewAddress(data.Default()).CreateAddress(ctx, *createAddressReq)
	if err != nil {
		resp.ErrCode = base.ErrCode_CreateAddressError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	addressInfo := new(user.AddressInfo)
	err = copier.Copy(addressInfo, info)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	resp.Info = addressInfo
	c.JSON(consts.StatusOK, resp)
}

// ShowAddress .
// @router /api/v1/user/addresses [GET]
func ShowAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	resp := new(user.ShowAddressResp)
	userID, err := getUserID(c)
	if err != nil {
		resp.ErrCode = base.ErrCode_UnauthorizedError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusUnauthorized, resp)
		return
	}
	info, err := logic.NewAddress(data.Default()).ShowAddress(ctx, userID)
	if err != nil {
		resp.ErrCode = base.ErrCode_ShowAddressError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	addressInfo := make([]*user.AddressInfo, len(info))
	err = copier.Copy(&addressInfo, &info)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	resp.Info = addressInfo
	c.JSON(consts.StatusOK, resp)
}

// UpdateAddress .
// @router /api/v1/user/addresses [PUT]
func UpdateAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateAddressReq
	resp := new(base.BaseResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	var updateAddressReq domain.UpdateAddressReq
	err = copier.Copy(&updateAddressReq, &req)
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
	updateAddressReq.UserID = userID
	err = logic.NewAddress(data.Default()).UpdateAddress(ctx, updateAddressReq)

	if err != nil {
		resp.ErrCode = base.ErrCode_UpdateAddressError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}

// DeleteAddress .
// @router /api/v1/user/addresses [DELETE]
func DeleteAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteAddressReq
	err = c.BindAndValidate(&req)
	resp := new(base.BaseResp)
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
	err = logic.NewAddress(data.Default()).DeleteAddress(ctx, req.Id, userID)

	if err != nil {
		resp.ErrCode = base.ErrCode_DeleteAddressError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}
