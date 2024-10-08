// Code generated by hertz generator.

package user

import (
	"context"

	"github.com/Tracy-coder/e-mall/biz/logic"
	base "github.com/Tracy-coder/e-mall/biz/model/base"
	user "github.com/Tracy-coder/e-mall/biz/model/user"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
)

// CreateFavourite .
// @router /api/v1/user/favourites [POST]
func CreateFavourite(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateFavouriteReq
	resp := new(user.CreateFavouriteResp)
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
	info, err := logic.NewFavourite(data.Default()).CreateFavourite(ctx, userID, req.ProductID)
	if err != nil {
		resp.ErrCode = base.ErrCode_CreateFavouriteError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	favouriteInfo := new(user.Favourite)
	err = copier.Copy(favouriteInfo, info)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	resp.Favourite = favouriteInfo
	c.JSON(consts.StatusOK, resp)
}

// ShowFavourites .
// @router /api/v1/user/favourites [GET]
func ShowFavourites(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ShowFavouritesReq
	resp := new(user.ShowFavouriteResp)
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
	info, total, err := logic.NewFavourite(data.Default()).ShowFavourite(ctx, userID, int(req.PageId), int(req.PageSize))
	if err != nil {
		resp.ErrCode = base.ErrCode_ShowFavouriteError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	favouriteInfo := make([]*user.Favourite, len(info))
	err = copier.Copy(&favouriteInfo, &info)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	resp.Favourites = favouriteInfo
	resp.Total = int32(total)
	c.JSON(consts.StatusOK, resp)
}

// DeleteFavourite .
// @router /api/v1/user/favourites [DELETE]
func DeleteFavourite(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteFavouriteReq
	resp := new(base.BaseResp)
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
	err = logic.NewFavourite(data.Default()).DeleteFavourite(ctx, userID, req.ProductID)
	if err != nil {
		resp.ErrCode = base.ErrCode_DeleteFavouriteError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusUnauthorized, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	c.JSON(consts.StatusOK, resp)
}
