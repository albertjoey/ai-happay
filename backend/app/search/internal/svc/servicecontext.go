package svc

import (
	"happy/app/search/internal/config"
	"happy/common/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	ES     *utils.ESClient
	DBSearch *utils.DBSearch
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	dsn := c.MySQL.User + ":" + c.MySQL.Password + "@tcp(" + c.MySQL.Host + ":3306)/" + c.MySQL.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 初始化Elasticsearch客户端
	var es *utils.ESClient
	es, err = utils.NewESClient(c.Elasticsearch.Hosts, c.Elasticsearch.Username, c.Elasticsearch.Password, "happy_content")
	if err != nil {
		// ES连接失败，使用数据库搜索
		es = nil
	}

	// 创建索引（如果不存在）
	if es != nil {
		if err := es.CreateIndex(); err != nil {
			// 索引可能已存在，忽略错误
		}
	}

	// 初始化数据库搜索
	dbSearch := utils.NewDBSearch(db)

	return &ServiceContext{
		Config:   c,
		DB:       db,
		ES:       es,
		DBSearch: dbSearch,
	}
}
