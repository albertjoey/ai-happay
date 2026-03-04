package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySQL struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	Redis struct {
		Host string
		Type string
	}
	// 热度更新配置
	HotScore struct {
		EnableCron bool   `json:",default=true"`  // 是否启用定时任务
		CronSpec   string `json:",default=0 * * * * *"` // 每小时更新一次
	} `json:",optional"`
}
