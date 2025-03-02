package lib

import (
	"fmt"
	"gorm.io/gorm"
)

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
