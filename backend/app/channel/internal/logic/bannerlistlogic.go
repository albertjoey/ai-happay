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
	// 使用Repository接口
	banners, err := l.svcCtx.BannerRepo.List(l.ctx, channelID)
	if err != nil {
		return nil, err
	}

	// 转换响应格式
	list := make([]BannerItem, 0, len(banners))
	for _, b := range banners {
		if status > 0 && b.Status != int8(status) {
			continue
		}
		list = append(list, BannerItem{
			ID:       b.ID,
			Title:    b.Title,
			Image:    b.Image,
			LinkType: int(parseLinkType(b.LinkType)),
			LinkURL:  b.LinkURL,
			Sort:     b.Sort,
		})
	}

	return &BannerListResponse{List: list}, nil
}

func parseLinkType(linkType string) int {
	switch linkType {
	case "material":
		return 1
	case "channel":
		return 2
	case "url":
		return 3
	default:
		return 0
	}
}
