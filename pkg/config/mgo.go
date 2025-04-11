package config

import (
	"context"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MgoConfig struct {
	URL        string
	DBName     string
	Collection string
	// 超时控制
	ctx context.Context
	// 设置对应的监听器
	m *event.CommandMonitor
}

type MgoConfigOption interface {
	Option(mg *MgoConfig)
}

type MgoConfigOptionFunc func(mg *MgoConfig)

func (fn MgoConfigOptionFunc) Option(mg *MgoConfig) {
	fn(mg)
}

func WithMgoURL(url string) MgoConfigOption {
	return MgoConfigOptionFunc(func(mg *MgoConfig) {
		mg.URL = url
	})
}

func WithMgDBName(name string) MgoConfigOption {
	return MgoConfigOptionFunc(func(mg *MgoConfig) {
		mg.DBName = name
	})
}

func WithMgCollection(name string) MgoConfigOption {
	return MgoConfigOptionFunc(func(mg *MgoConfig) {
		mg.Collection = name
	})
}

func WithContext(ctx context.Context) MgoConfigOption {
	return MgoConfigOptionFunc(func(mg *MgoConfig) {
		mg.ctx = ctx
	})
}

func NewMgoConfig(options ...MgoConfigOption) *MgoConfig {
	mgo := &MgoConfig{}

	for _, opt := range options {
		opt.Option(mgo)
	}

	return mgo
}

func NewMgo(config *MgoConfig) (*mongo.Database, error) {
	client, err := mongo.Connect(
		config.ctx,
		options.Client().ApplyURI(config.URL))

	if err != nil {
		return nil, err
	}

	return client.Database(config.DBName), nil
}
