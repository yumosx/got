package lv_reflect

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/yumosx/oneframe/lv_cache"
)

func LoadFileByLocale(locale string) error {
	var dataMap map[string]any
	bytes, err := os.ReadFile("locales/" + locale + ".json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &dataMap)
	if err != nil {
		return err
	}
	lv_cache.GetCacheClient().HMSet(context.Background(), locale, dataMap, time.Hour)
	return err
}

func GetTextLocale(localKey, key string) string {
	v, _ := lv_cache.GetCacheClient().HGet(context.Background(), localKey, key)
	if v == "" {
		v = key
	}
	return v
}
