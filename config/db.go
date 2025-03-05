package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	UserName string
	Password string
	Port     string
	DBName   string
}

type DBOption interface{
	apply(config DBConfig) DBConfig
}

type OptionFunc func(config DBConfig) DBConfig

func (o OptionFunc) apply(config DBConfig) {
	o(config)
}

func WithUserName(name string) OptionFunc {
	return OptionFunc(func(config DBConfig) DBConfig {
		config.UserName = name
		return config
	})
}

func WithPassword(password string) OptionFunc {
	return OptionFunc(func(config DBConfig) DBConfig {
		config.Password = password
		return config
	})
}

func WithPort(port string) OptionFunc {
	return OptionFunc(func(config DBConfig) DBConfig {
		config.Port = port
		return config
	})
}

func WithDBName(dbname string) OptionFunc {
	return OptionFunc(func(config DBConfig) DBConfig {
		config.DBName = dbname
		return config
	})
}

// InitDB 根据参数建立对应的数据库连接
func InitDB(username, password, host, port, dbname string) *gorm.DB {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
	)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func TearTables(db *gorm.DB, tables ...string) error {
	//1. 遍历所有的 tables, 然后打印
	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s", table)
		err := db.Exec(query).Error
		if err != nil {
			return err
		}
	}

	return nil
}
