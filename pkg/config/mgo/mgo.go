package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URL        string
	DBName     string
	Collection string
	// 超时控制
	ctx context.Context
	// 设置对应的监听器
	m *event.CommandMonitor
}

type ConfigOption interface {
	Option(mg *Config)
}

type ConfigOptionFunc func(mg *Config)

func (fn ConfigOptionFunc) Option(mg *Config) {
	fn(mg)
}

func WithMgoURL(url string) ConfigOption {
	return ConfigOptionFunc(func(mg *Config) {
		mg.URL = url
	})
}

func WithMgDBName(name string) ConfigOption {
	return ConfigOptionFunc(func(mg *Config) {
		mg.DBName = name
	})
}

func WithMgCollection(name string) ConfigOption {
	return ConfigOptionFunc(func(mg *Config) {
		mg.Collection = name
	})
}

func WithContext(ctx context.Context) ConfigOption {
	return ConfigOptionFunc(func(mg *Config) {
		mg.ctx = ctx
	})
}

func NewMgoConfig(options ...ConfigOption) *Config {
	mgo := &Config{}

	for _, opt := range options {
		opt.Option(mgo)
	}

	return mgo
}

func NewMgo(config *Config) (*mongo.Database, error) {
	client, err := mongo.Connect(
		config.ctx,
		options.Client().ApplyURI(config.URL))

	if err != nil {
		return nil, err
	}

	return client.Database(config.DBName), nil
}
