package cache

import (
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"
)

type CacheSuite struct {
	suite.Suite
	rdb redis.Cmdable
}

func TestCache(t *testing.T) {
	suite.Run(t, new(CacheSuite))
}

func (s *CacheSuite) SetupSuite() {
	configData, err := os.ReadFile("test_config.yaml")
	if err != nil {
		s.T().Fatalf("Failed to read config file: %v", err)
	}
	var config struct {
		Redis struct {
			Addr     string `yaml:"addr"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
		} `yaml:"redis"`
	}

	if err := yaml.Unmarshal(configData, &config); err != nil {
		s.T().Fatalf("Failed to parse config file: %v", err)
	}

	s.rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
}

func (s *CacheSuite) TestSet() {
	ctx := s.T().Context()
	cache := NewCache(s.rdb)
	err := cache.Set(ctx, "test_set_key", "value", 0)
	require.NoError(s.T(), err)
	result, err := s.rdb.Get(ctx, "test_set_key").Result()
	require.NoError(s.T(), err)
	require.Equal(s.T(), "value", result)
}

func (s *CacheSuite) TearDownSuite() {
	ctx := s.T().Context()
	err := s.rdb.FlushDB(ctx).Err()
	if err != nil {
		s.T().Logf("Warning: Failed to flush Redis database: %v", err)
	}
}

func (s *CacheSuite) TestSetNX() {
	t := s.T()
	t.Run("SetNX_KeyNotExists", func(t *testing.T) {
		ctx := t.Context()
		cache := NewCache(s.rdb)
		b, err := cache.SetNX(ctx, "test_setnx_not_exists_key", "value", 0)
		require.NoError(t, err)
		assert.True(t, b)
	})

	t.Run("SetNX_KeyExists", func(t *testing.T) {
		ctx := t.Context()
		cache := NewCache(s.rdb)
		s.rdb.Set(ctx, "test_setnx_exists_key", "existing_value", 0)
		b, err := cache.SetNX(ctx, "test_setnx_exists_key", "value", 0)
		require.NoError(t, err)
		assert.False(t, b)
	})
}
