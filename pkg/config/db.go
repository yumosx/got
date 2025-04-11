package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConfig 数据库信息配置
// 1. 通过 function option 模式配置
// 2. 通过 toml 文件进行配置
type DBConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
}

type DBConfigOption interface {
	Option(config *DBConfig)
}

type DBConfigOptionFunc func(config *DBConfig)

func (fn DBConfigOptionFunc) Option(config *DBConfig) {
	fn(config)
}

func WithUserName(userName string) DBConfigOptionFunc {
	return DBConfigOptionFunc(func(config *DBConfig) {
		config.UserName = userName
	})
}

func WithPassword(password string) DBConfigOptionFunc {
	return DBConfigOptionFunc(func(config *DBConfig) {
		config.Password = password
	})
}

func WithHost(host string) DBConfigOptionFunc {
	return DBConfigOptionFunc(func(config *DBConfig) {
		config.Host = host
	})
}

func WithPort(port string) DBConfigOptionFunc {
	return DBConfigOptionFunc(func(config *DBConfig) {
		config.Port = port
	})
}

func WithDBName(db string) DBConfigOptionFunc {
	return DBConfigOptionFunc(func(config *DBConfig) {
		config.DBName = db
	})
}

func NewConfig(options ...DBConfigOption) *DBConfig {
	config := &DBConfig{}
	for _, opt := range options {
		opt.Option(config)
	}

	return config
}

// NewDB 根据参数建立对应的数据库连接
func NewDB(config *DBConfig) (*gorm.DB, error) {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

func TearTables(db *gorm.DB, tables ...string) error {
	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s", table)
		err := db.Exec(query).Error
		if err != nil {
			return err
		}
	}
	return nil
}
