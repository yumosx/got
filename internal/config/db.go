package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
