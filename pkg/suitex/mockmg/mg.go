package mockmg

import (
	"github.com/mjarkk/mongomock"
)

type MongoMockAdapter struct {
	db *mongomock.TestConnection
}

func NewMongoMockAdapter() *MongoMockAdapter {
	return &MongoMockAdapter{db: mongomock.NewDB()}
}
