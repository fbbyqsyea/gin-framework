package global

import (
	"github.com/fbbyqsyea/gin-framework/config"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// 全局变量定义
var (
	CONFIG config.Config // 配置
	LOGGER *zap.Logger   // 日志模块
	DB     struct {      // 数据库
		Master  *sqlx.DB
		Replica *sqlx.DB
	}
	REDIS *redis.Client // redis
)
