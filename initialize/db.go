package initialize

import (
	"fmt"

	"github.com/fbbyqsyea/gin-framework/config"
	"github.com/fbbyqsyea/gin-framework/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 初始化db 主从配置
func initDB() {
	// 主库
	global.DB.Master, _ = newDb(global.CONFIG.DB.Master)
	// 从库
	global.DB.Replica, _ = newDb(global.CONFIG.DB.Replica)
}

// 数据库连接
func newDb(cfg config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect(cfg.Engine, cfg.Dsn())
	if err != nil {
		global.LOGGER.Error(err.Error())
		panic(fmt.Errorf("fatal error connect %s: [%s]", cfg.Engine, err.Error()))
	}
	return db, nil
}
