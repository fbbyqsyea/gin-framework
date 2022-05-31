package initialize

import (
	"context"
	"fmt"

	"github.com/fbbyqsyea/gin-framework/config"
	"github.com/fbbyqsyea/gin-framework/global"
	"github.com/go-redis/redis/v8"
)

// 初始化redis缓存 单机配置
func initRedis() {
	global.REDIS, _ = newRedis(global.CONFIG.Redis, true)
}

// redis连接
func newRedis(c config.Redis, isMust bool) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		if isMust {
			panic(fmt.Errorf("fatal error ping: [%s]", err))
		} else {
			return nil, err
		}
	}
	return client, nil
}
