// Package redis 工具包
package redis

import (
	"BAT-douyin/setting"
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	*redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次

var Redis *RedisClient

func Init(conf *setting.RedisConfig) error {

	// 初始化自定的 RedisClient 实例
	Redis = &RedisClient{}
	// 使用默认的 context
	Redis.Context = context.Background()

	// 使用 redis 库里的 NewClient 初始化连接
	Redis.Client = redis.NewClient(&redis.Options{
		//Addr: fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Addr: "127.0.0.1:6379",
		DB:   conf.DB,
	})

	// 测试一下连接
	err := Redis.Ping()
	if err != nil {
		return err
	}
	return nil

}

// Ping 用以测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			return ""
		}
	}
	return result
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		return false
	}
	return true
}
