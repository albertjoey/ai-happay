package svc

import (
	"fmt"
	"happy/app/recommend/internal/config"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Cron   *cron.Cron
}

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

	ctx := &ServiceContext{
		Config: c,
		DB:     db,
		Cron:   cron.New(cron.WithSeconds()),
	}

	// 启动热度更新定时任务
	if c.HotScore.EnableCron {
		ctx.startHotScoreUpdateTask()
	}

	return ctx
}

// startHotScoreUpdateTask 启动热度分数更新定时任务
func (s *ServiceContext) startHotScoreUpdateTask() {
	cronSpec := s.Config.HotScore.CronSpec
	if cronSpec == "" {
		cronSpec = "0 0 * * * *" // 默认每小时执行一次
	}

	_, err := s.Cron.AddFunc(cronSpec, func() {
		logx.Info("开始更新热度分数...")
		startTime := time.Now()
		
		// 批量更新热度分数
		err := s.UpdateAllHotScores()
		if err != nil {
			logx.Errorf("更新热度分数失败: %v", err)
			return
		}
		
		logx.Infof("热度分数更新完成，耗时: %v", time.Since(startTime))
	})

	if err != nil {
		logx.Errorf("添加定时任务失败: %v", err)
		return
	}

	s.Cron.Start()
	logx.Info("热度更新定时任务已启动")
}

// UpdateAllHotScores 更新所有内容的热度分数
func (s *ServiceContext) UpdateAllHotScores() error {
	// 使用SQL批量计算和更新热度分数
	// 热度分数 = (浏览×1 + 点赞×5 + 评论×10 + 收藏×8 + 分享×15) / (时间衰减)
	// 时间衰减 = (小时数 + 1)^0.5
	updateSQL := `
		UPDATE content 
		SET hot_score = (
			(view_count * 1 + like_count * 5 + comment_count * 10 + collect_count * 8 + share_count * 15) / 
			POW(TIMESTAMPDIFF(HOUR, COALESCE(publish_at, created_at), NOW()) + 1, 0.5)
		)
		WHERE deleted_at IS NULL AND status = 1
	`
	
	result := s.DB.Exec(updateSQL)
	if result.Error != nil {
		return result.Error
	}
	
	logx.Infof("更新了 %d 条内容的热度分数", result.RowsAffected)
	return nil
}

// UpdateHotScore 手动更新单个内容的热度分数
func (s *ServiceContext) UpdateHotScore(contentID uint) error {
	updateSQL := `
		UPDATE content 
		SET hot_score = (
			(view_count * 1 + like_count * 5 + comment_count * 10 + collect_count * 8 + share_count * 15) / 
			POW(TIMESTAMPDIFF(HOUR, COALESCE(publish_at, created_at), NOW()) + 1, 0.5)
		)
		WHERE id = ? AND deleted_at IS NULL
	`
	
	return s.DB.Exec(updateSQL, contentID).Error
}
