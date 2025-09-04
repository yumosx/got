package ormx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/yumosx/got/pkg/config/db"
	"gorm.io/gorm"
)

type BatchExportSuite struct {
	suite.Suite
	db *gorm.DB
}

func NewBatchExportSuite() *BatchExportSuite {
	return &BatchExportSuite{}
}

func (b *BatchExportSuite) TearDown() {
	err := db.TearTables(b.db, "")
	require.NoError(b.T(), err)
}

func (b *BatchExportSuite) SetupTest() {
	t := b.T()
	config := db.NewConfig(db.WithPort("3306"))
	gdb, err := db.NewDB(config)
	require.NoError(t, err)
	b.db = gdb
}

func (b *BatchExportSuite) TestBatchExport() {
	type Model struct {
		Id   string `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}

	testcases := []struct {
		name   string
		before func()
	}{
		{
			name: "批量导出数据",
			before: func() {
			},
		},
	}

	t := b.T()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			batchConfig := NewBatchConfig(2, b.db, WithBatch(1000), WithMaxRecords(1000))
			_, err := BatchExport[Model](batchConfig, func(offset, currentLimit int, db *gorm.DB) ([]Model, error) {
				return nil, errors.New("")
			})
			require.NoError(t, err)
		})
	}
}
