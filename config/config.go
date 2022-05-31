package config

import (
	"fmt"
	"time"
)

// 配置信息
type Config struct {
	System System `json:"core" yaml:"system"` // 系统配置
	DB     DB     `json:"db" yaml:"db"`
	Redis  Redis  `json:"redis" yaml:"redis"`
	Logger Logger `json:"logger" yaml:"logger"`
	Jwt    JWT    `json:"jwt" yaml:"jwt"`
}

// 系统配置
type System struct {
	Mode         string        `json:"env" yaml:"mode"`                  // 运行模式
	Host         string        `json:"host" yaml:"host"`                 // 地址
	Port         int           `json:"port" yaml:"port"`                 // 端口号
	ReadTimeout  time.Duration `json:"readTimeOut" yaml:"readTimeOut"`   // 读取超时时间
	WriteTimeout time.Duration `json:"writeTimeout" yaml:"writeTimeout"` // 写入超时时间
}

func (c *System) Attr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// 数据库配置
type DB struct {
	Master  DBConfig `json:"master" yaml:"master"`
	Replica DBConfig `json:"replica" yaml:"replica"`
}

type DBConfig struct {
	Engine   string `json:"engine" yaml:"engine"`     // 数据库引擎
	Host     string `json:"host" yaml:"host"`         // 服务器地址
	Port     string `json:"port" yaml:"port"`         // 服务器端口
	UserName string `json:"username" yaml:"username"` // MySQL用户名
	Password string `json:"password" yaml:"password"` // mysql密码
	DbName   string `json:"dbname" yaml:"dbname"`     // 数据库
	Config   string `json:"config" yaml:"config"`     // 额外配置信息
}

// db dsn
func (db *DBConfig) Dsn() (dsn string) {
	switch db.Engine {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", db.UserName, db.Password, db.Host, db.Port, db.DbName, db.Config)
	}
	return
}

// redis缓存配置
type Redis struct {
	DB       int    `json:"db" yaml:"db"`             // 数据库
	Addr     string `json:"addr" yaml:"addr"`         // 地址 ip:port
	Password string `json:"password" yaml:"password"` // 密码
}

// 配置模块配置
type Logger struct {
	Format        string `json:"format" yaml:"format"`
	Director      string `json:"director" yaml:"director"`
	ShowLine      bool   `json:"showLine" yaml:"showLine"`
	EncodeLevel   string `json:"encodeLevel" yaml:"encodeLevel"`
	StackTraceKey string `json:"stackTraceKey" yaml:"stackTraceKey"`
}

// jtw配置
type JWT struct {
	SecretKey  string `json:"secretKey" yaml:"secretKey"`
	ExpireTime int64  `json:"expireTime" yaml:"expireTime"`
	Issuer     string `json:"issuer" yaml:"issuer"`
	Subject    string `json:"subject" yaml:"subject"`
}
