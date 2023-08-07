package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang-stagging/core"
	"time"
)

// InitRedis 初始化
func InitRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", core.Config.Redis.IP, core.Config.Redis.Port),
		Password: core.Config.Redis.Password,
		DB:       core.Config.Redis.Db,
		PoolSize: core.Config.Redis.PoolSize,
	})
	redis.SetLogger(&RDBLogger{Logger: core.Logger})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	return rdb, err
}
