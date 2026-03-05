package svc

import (
	"fmt"
	"happy/app/channel/internal/config"
	"happy/app/channel/internal/repository"
	"happy/common/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ServiceContext 服务上下文
// 采用依赖注入方式，注入Repository接口而非具体实现
type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB

	// Repository接口 - 依赖倒置原则
	ChannelRepo     repository.ChannelRepository
	MaterialRepo    repository.MaterialRepository
	ChapterRepo     repository.ChapterRepository
	RoleRepo        repository.RoleRepository
	PermissionRepo  repository.PermissionRepository
	AdminUserRepo   repository.AdminUserRepository
	BannerRepo      repository.BannerRepository
	DiamondRepo     repository.DiamondRepository
	RecommendRepo   repository.RecommendRepository
	AdSlotRepo      repository.AdSlotRepository
	FeedConfigRepo  repository.FeedConfigRepository
	InteractionRepo repository.InteractionRepository
}

// NewServiceContext 创建服务上下文
func NewServiceContext(c config.Config) *ServiceContext {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local",
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	// 设置字符集
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// 执行SET NAMES utf8mb4
	db.Exec("SET NAMES utf8mb4")

	// 自动迁移
	db.AutoMigrate(&model.Channel{})

	// 创建Repository实例 - 依赖注入
	return &ServiceContext{
		Config: c,
		DB:     db,

		// 注入Repository实现
		ChannelRepo:     repository.NewChannelRepository(db),
		MaterialRepo:    repository.NewMaterialRepository(db),
		ChapterRepo:     repository.NewChapterRepository(db),
		RoleRepo:        repository.NewRoleRepository(db),
		PermissionRepo:  repository.NewPermissionRepository(db),
		AdminUserRepo:   repository.NewAdminUserRepository(db),
		BannerRepo:      repository.NewBannerRepository(db),
		DiamondRepo:     repository.NewDiamondRepository(db),
		RecommendRepo:   repository.NewRecommendRepository(db),
		AdSlotRepo:      repository.NewAdSlotRepository(db),
		FeedConfigRepo:  repository.NewFeedConfigRepository(db),
		InteractionRepo: repository.NewInteractionRepository(db),
	}
}
