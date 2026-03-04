package logic

import (
	"context"

	"happy/app/channel/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BannerListLogic {
	return &BannerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type BannerItem struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	LinkType int    `json:"link_type"`
	LinkURL  string `json:"link_url"`
	Sort     int    `json:"sort"`
}

type BannerListResponse struct {
	List []BannerItem `json:"list"`
}

func (l *BannerListLogic) BannerList(channelID uint, status int) (*BannerListResponse, error) {
	query := `
		SELECT id, title, image, link_type, link_url, sort
		FROM banner
		WHERE deleted_at IS NULL
	`
	args := []interface{}{}

	if channelID > 0 {
		query += " AND channel_id = ?"
		args = append(args, channelID)
	}

	if status > 0 {
		query += " AND status = ?"
		args = append(args, status)
	}

	query += " ORDER BY sort ASC"

	var banners []BannerItem
	err := l.svcCtx.DB.Raw(query, args...).Scan(&banners).Error
	if err != nil {
		return nil, err
	}

	if banners == nil {
		banners = []BannerItem{}
	}

	return &BannerListResponse{List: banners}, nil
}
