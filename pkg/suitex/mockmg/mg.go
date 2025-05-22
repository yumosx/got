package mockmg

import (
	"github.com/mjarkk/mongomock"
)

type MongoMockAdapter struct {
	DB *mongomock.TestConnection
}

func NewMongoMockAdapter() *MongoMockAdapter {
	return &MongoMockAdapter{DB: mongomock.NewDB()}
}
