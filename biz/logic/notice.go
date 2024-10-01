package logic

import (
	"context"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/jinzhu/copier"
)

type Notice struct {
	Data *data.Data
}

func NewNotice(data *data.Data) domain.Notice {
	return &Notice{
		Data: data,
	}
}

func (c *Notice) CreateNotice(ctx context.Context, text string) (*domain.NoticeInfo, error) {
	noticeDB, err := c.Data.DBClient.Notice.Create().
		SetText(text).Save(ctx)
	if err != nil {
		return nil, err
	}
	notice := new(domain.NoticeInfo)
	copier.Copy(notice, noticeDB)
	if err != nil {
		return nil, err
	}
	notice.CreatedAt = uint64(noticeDB.CreatedAt.Unix())
	return notice, nil
}

func (c *Notice) ShowNotice(ctx context.Context, page_id int, page_size int) ([]*domain.NoticeInfo, int, error) {
	total := c.Data.DBClient.Notice.Query().CountX(ctx)
	noticesDB, err := c.Data.DBClient.Notice.Query().Offset(int(page_size) * (int(page_id - 1))).Limit(int(page_size)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	info := make([]*domain.NoticeInfo, len(noticesDB))
	for i, v := range noticesDB {
		info[i] = &domain.NoticeInfo{
			ID:        v.ID,
			Text:      v.Text,
			CreatedAt: uint64(v.CreatedAt.Unix()),
		}
	}

	return info, total, nil
}
