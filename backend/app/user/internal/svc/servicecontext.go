package svc

import (
	"happy/app/user/internal/config"
	"happy/common/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	dsn := c.MySQL.User + ":" + c.MySQL.Password + "@tcp(" + c.MySQL.Host + ":3306)/" + c.MySQL.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动迁移
	db.AutoMigrate(
		&model.User{},
		&model.Tenant{},
		&model.Follow{},
	)

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
