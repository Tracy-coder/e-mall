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

// ShowNotices .
// @router /api/v1/notices [GET]
func ShowNotices(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ShowNoticesReq
	resp := new(user.ShowNoticesResp)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.ErrCode = base.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	info, total, err := logic.NewNotice(data.Default()).ShowNotice(ctx, int(req.PageID), int(req.PageSize))
	if err != nil {
		resp.ErrCode = base.ErrCode_ShowNoticesError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	noticeInfo := make([]*user.NoticeInfo, len(info))
	err = copier.Copy(&noticeInfo, &info)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()

		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	resp.Notices = noticeInfo
	resp.Total = int32(total)
	c.JSON(consts.StatusOK, resp)
}

// CreateNotice .
// @router /api/v2/notices [POST]
func CreateNotice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateNoticeReq
	err = c.BindAndValidate(&req)
	resp := new(user.CreateNoticeResp)
	if err != nil {
		resp.ErrCode = base.ErrCode_ArgumentError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	info, err := logic.NewNotice(data.Default()).CreateNotice(ctx, req.Text)
	if err != nil {
		resp.ErrCode = base.ErrCode_CreateNoticeError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	noticeInfo := new(user.NoticeInfo)
	err = copier.Copy(noticeInfo, info)
	if err != nil {
		resp.ErrCode = base.ErrCode_CopierError
		resp.ErrMsg = err.Error()
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	resp.ErrCode = base.ErrCode_Success
	resp.ErrMsg = "success"
	resp.Notice = noticeInfo
	c.JSON(consts.StatusOK, resp)
}
