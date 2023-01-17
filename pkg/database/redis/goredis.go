package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"time"
)

type GoRedisClient struct {
	driver *redis.Client
}

type GoRedisResult struct {
	result string
	err    error
}

func (goRedisResult *GoRedisResult) Result() (string, error) {
	return goRedisResult.result, goRedisResult.err
}

func (goRedisResult *GoRedisResult) String() string {
	return goRedisResult.result
}

// Connect 连接
func (client *GoRedisClient) Connect() ClientInterface {
	addr := os.Getenv("redis.host") + ":" + os.Getenv("redis.port")
	db, _ := strconv.Atoi(os.Getenv("redis.db"))
	client.driver = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("redis.password"),
		DB:       db,
	})
	return client
}

// Set 设置缓存
func (client *GoRedisClient) Set(key string, value interface{}, duration time.Duration) CmdResultInterface {
	result, err := client.driver.Set(context.Background(), key, value, duration).Result()
	return &GoRedisResult{result: result, err: err}
}

// Get 获取缓存
func (client *GoRedisClient) Get(key string) CmdResultInterface {
	result, err := client.driver.Get(context.Background(), key).Result()
	return &GoRedisResult{result: result, err: err}
}

// Del 删除缓存
func (client *GoRedisClient) Del(key string) error {
	_, err := client.driver.Del(context.Background(), key).Result()
	return err
}

// FlushDB 清除所有
func (client *GoRedisClient) FlushDB() error {
	return client.driver.FlushDB(context.Background()).Err()
}

// Close 关闭连接
func (client *GoRedisClient) Close() error {
	return client.driver.Close()
}

func (client *GoRedisClient) SetNx(key string, value interface{}, duration time.Duration) (bool, error) {
	return client.driver.SetNX(context.Background(), key, value, duration).Result()
}

func (client *GoRedisClient) Increment(key string, duration time.Duration) (int64, error) {
	result, err := client.driver.Incr(context.Background(), key).Result()
	client.driver.Expire(context.Background(), key, duration)
	return result, err
}

func (client *GoRedisClient) Decrement(key string, duration time.Duration) (int64, error) {
	result, err := client.driver.Decr(context.Background(), key).Result()
	client.driver.Expire(context.Background(), key, duration)
	return result, err
}

func (client *GoRedisClient) LPop(key string) CmdResultInterface {
	length, err := client.driver.LLen(context.Background(), key).Result()
	if err != nil || length <= 0 {
		return &GoRedisResult{result: "", err: fmt.Errorf("not found")}
	}
	result, err := client.driver.LPop(context.Background(), key).Result()
	return &GoRedisResult{result: result, err: err}
}

func (client *GoRedisClient) RPush(key string, value interface{}) CmdResultInterface {
	client.Connect()
	defer client.Close()
	_, err := client.driver.RPush(context.Background(), key, value).Result()
	return &GoRedisResult{result: "", err: err}
}
