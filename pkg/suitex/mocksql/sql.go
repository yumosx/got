package mocksql

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLMockAdapter struct {
	DB   *gorm.DB
	Mock sqlmock.Sqlmock
}

func NewSQLAdapter() (*SQLMockAdapter, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	d := mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	})
	gdb, err := gorm.Open(d, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {

	}
	return &SQLMockAdapter{
		DB:   gdb,
		Mock: mock,
	}, nil
}

func (adapter *SQLMockAdapter) AddRows() {

}
