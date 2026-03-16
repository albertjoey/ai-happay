package main

import (
	"flag"
	"fmt"
	"net/http"

	"happy/app/channel/internal/config"
	"happy/app/channel/internal/handler"
	"happy/app/channel/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/channel/etc/channel.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	// 设置响应头,确保UTF-8编码
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			next(w, r)
		}
	})

	ctx := svc.NewServiceContext(c)

	// 注册路由
	registerRoutes(server, ctx)

	fmt.Printf("Starting channel server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func registerRoutes(server *rest.Server, ctx *svc.ServiceContext) {
	// 频道列表
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/channel/list",
			Handler: handler.ChannelListHandler(ctx),
		},
	)

	// 创建频道
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/channel",
			Handler: handler.ChannelCreateHandler(ctx),
		},
	)

	// 更新频道
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/channel/:id",
			Handler: handler.ChannelUpdateHandler(ctx),
		},
	)

	// 删除频道
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/channel/:id",
			Handler: handler.ChannelDeleteHandler(ctx),
		},
	)

	// 频道排序
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/channel/sort",
			Handler: handler.ChannelSortHandler(ctx),
		},
	)

	// 获取频道配置
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/channel/config/:id",
			Handler: handler.ChannelConfigGetHandler(ctx),
		},
	)

	// 更新频道配置
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/channel/config/:id",
			Handler: handler.ChannelConfigUpdateHandler(ctx),
		},
	)

	// ==================== 金刚位路由 ====================
	// 金刚位列表
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/diamond/list",
			Handler: handler.DiamondListHandler(ctx),
		},
	)

	// 创建金刚位
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/diamond",
			Handler: handler.DiamondCreateHandler(ctx),
		},
	)

	// 更新金刚位
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/diamond/:id",
			Handler: handler.DiamondUpdateHandler(ctx),
		},
	)

	// 删除金刚位
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/diamond/:id",
			Handler: handler.DiamondDeleteHandler(ctx),
		},
	)

	// ==================== 推荐位路由 ====================
	// 推荐位列表
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/recommend/list",
			Handler: handler.RecommendListHandler(ctx),
		},
	)

	// 创建推荐位
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/recommend",
			Handler: handler.RecommendCreateHandler(ctx),
		},
	)

	// 更新推荐位
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/recommend/:id",
			Handler: handler.RecommendUpdateHandler(ctx),
		},
	)

	// 删除推荐位
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/recommend/:id",
			Handler: handler.RecommendDeleteHandler(ctx),
		},
	)

	// ==================== Feed流路由 ====================
	// Feed流配置列表
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/feed-config/list",
			Handler: handler.FeedConfigListHandler(ctx),
		},
	)

	// 创建Feed流配置
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/feed-config",
			Handler: handler.FeedConfigCreateHandler(ctx),
		},
	)

	// 更新Feed流配置
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/feed-config/:id",
			Handler: handler.FeedConfigUpdateHandler(ctx),
		},
	)

	// 删除Feed流配置
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/feed-config/:id",
			Handler: handler.FeedConfigDeleteHandler(ctx),
		},
	)

	// ==================== 广告位路由 ====================
	// 广告位列表
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/ad-slot/list",
			Handler: handler.AdSlotListHandler(ctx),
		},
	)

	// 创建广告位
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/ad-slot",
			Handler: handler.AdSlotCreateHandler(ctx),
		},
	)

	// 更新广告位
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/ad-slot/:id",
			Handler: handler.AdSlotUpdateHandler(ctx),
		},
	)

	// 删除广告位
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/ad-slot/:id",
			Handler: handler.AdSlotDeleteHandler(ctx),
		},
	)

	// ==================== 物料管理路由 ====================
	// 物料列表
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/material/list",
			Handler: handler.MaterialListHandler(ctx),
		},
	)

	// 物料详情
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/material/:id",
			Handler: handler.MaterialDetailHandler(ctx),
		},
	)

	// 创建物料
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/material",
			Handler: handler.MaterialCreateHandler(ctx),
		},
	)

	// 更新物料
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/material/:id",
			Handler: handler.MaterialUpdateHandler(ctx),
		},
	)

	// 删除物料
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/material/:id",
			Handler: handler.MaterialDeleteHandler(ctx),
		},
	)

	// ==================== 推荐预览路由 ====================
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/recommend/preview",
			Handler: handler.RecommendPreviewHandler(ctx),
		},
	)

	// ==================== 章节路由 ====================
	// 章节列表
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/chapter/list/:id",
			Handler: handler.ChapterListHandler(ctx),
		},
	)

	// 章节详情
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/chapter/:id",
			Handler: handler.ChapterDetailHandler(ctx),
		},
	)

	// 章节创建
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/chapter",
			Handler: handler.ChapterCreateHandler(ctx),
		},
	)

	// 章节更新
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/chapter/:id",
			Handler: handler.ChapterUpdateHandler(ctx),
		},
	)

	// 章节删除
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/chapter/:id",
			Handler: handler.ChapterDeleteHandler(ctx),
		},
	)

	// ==================== Banner路由 ====================
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/banner/list",
			Handler: handler.BannerListHandler(ctx),
		},
	)

	// ==================== RBAC权限管理路由 ====================
	// 角色管理
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/role/list",
			Handler: handler.RoleListHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/role",
			Handler: handler.RoleCreateHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/role/:id",
			Handler: handler.RoleUpdateHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/role/:id",
			Handler: handler.RoleDeleteHandler(ctx),
		},
	)
	// 角色权限分配
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/role/:id/permissions",
			Handler: handler.RolePermissionsHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/role/:id/permissions",
			Handler: handler.AssignRolePermissionsHandler(ctx),
		},
	)

	// 权限管理
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/permission/tree",
			Handler: handler.PermissionTreeHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/permission",
			Handler: handler.PermissionCreateHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/permission/:id",
			Handler: handler.PermissionUpdateHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/permission/:id",
			Handler: handler.PermissionDeleteHandler(ctx),
		},
	)

	// 管理员管理
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/admin-user/list",
			Handler: handler.AdminUserListHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/admin-user",
			Handler: handler.AdminUserCreateHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/admin-user/:id",
			Handler: handler.AdminUserUpdateHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/admin-user/:id",
			Handler: handler.AdminUserDeleteHandler(ctx),
		},
	)
	// 管理员角色分配
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/admin-user/:id/roles",
			Handler: handler.AdminRolesHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/admin-user/:id/roles",
			Handler: handler.AssignAdminRolesHandler(ctx),
		},
	)

	// ==================== 互动系统路由 ====================
	// 点赞
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/like",
			Handler: handler.LikeHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/like/status",
			Handler: handler.LikeStatusHandler(ctx),
		},
	)

	// 收藏
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/collect",
			Handler: handler.CollectHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/collect/status",
			Handler: handler.CollectStatusHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/collect/list",
			Handler: handler.CollectListHandler(ctx),
		},
	)

	// 评论
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/comment",
			Handler: handler.CommentCreateHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/comment/list",
			Handler: handler.CommentListHandler(ctx),
		},
	)
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/comment/:id",
			Handler: handler.CommentDeleteHandler(ctx),
		},
	)

	// ==================== 话题管理路由 ====================
	// 话题列表
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/topic/list",
			Handler: handler.TopicListHandler(ctx),
		},
	)
	// 创建话题
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/topic",
			Handler: handler.TopicCreateHandler(ctx),
		},
	)
	// 更新话题
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/topic/:id",
			Handler: handler.TopicUpdateHandler(ctx),
		},
	)
	// 删除话题
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/topic/:id",
			Handler: handler.TopicDeleteHandler(ctx),
		},
	)

	// ==================== 标签管理路由 ====================
	// 标签列表
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/tag/list",
			Handler: handler.TagListHandler(ctx),
		},
	)
	// 创建标签
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/tag",
			Handler: handler.TagCreateHandler(ctx),
		},
	)
	// 更新标签
	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/tag/:id",
			Handler: handler.TagUpdateHandler(ctx),
		},
	)
	// 删除标签
	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/tag/:id",
			Handler: handler.TagDeleteHandler(ctx),
		},
	)

	// ==================== 发现页路由 ====================
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/discover",
			Handler: handler.DiscoverHandler(ctx),
		},
	)

	// 发现页模块配置
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/discover/config",
			Handler: handler.DiscoverConfigListHandler(ctx),
		},
	)

	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/discover/config",
			Handler: handler.DiscoverConfigUpdateHandler(ctx),
		},
	)

	// 发现页内容管理
	server.AddRoute(
		rest.Route{
			Method:  http.MethodGet,
			Path:    "/api/v1/discover/items",
			Handler: handler.DiscoverItemListHandler(ctx),
		},
	)

	server.AddRoute(
		rest.Route{
			Method:  http.MethodPost,
			Path:    "/api/v1/discover/items",
			Handler: handler.DiscoverItemCreateHandler(ctx),
		},
	)

	server.AddRoute(
		rest.Route{
			Method:  http.MethodPut,
			Path:    "/api/v1/discover/items",
			Handler: handler.DiscoverItemUpdateHandler(ctx),
		},
	)

	server.AddRoute(
		rest.Route{
			Method:  http.MethodDelete,
			Path:    "/api/v1/discover/items/:id",
			Handler: handler.DiscoverItemDeleteHandler(ctx),
		},
	)
}
