package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	// 数据库配置
	MySQL struct {
		Host         string
		Port         int
		User         string
		Password     string
		DBName       string
		MaxOpenConns int
		MaxIdleConns int
	}

	// Redis配置
	Redis cache.CacheConf

	// Elasticsearch配置
	Elasticsearch struct {
		Hosts    []string
		Username string
		Password string
	}

	// Kafka配置
	Kafka struct {
		Brokers []string
	}

	// JWT配置
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	// 租户配置
	Tenant struct {
		Enabled bool
	}
}
