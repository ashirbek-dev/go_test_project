package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

var KvDb KvRedis

type KvRedis struct {
	rc  *redis.Client
	ctx context.Context
}

func InitKvRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	rc := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	ctx := context.Background()
	KvDb = KvRedis{rc: rc, ctx: ctx}
}

// FIFO

func (n KvRedis) Push(key string, data any) {
	err := n.rc.RPush(n.ctx, key, data).Err()
	if err != nil {
		return
	}
}
func (n KvRedis) PushTask(key string, task any) error {
	var err error
	var taskData []byte
	taskData, err = json.Marshal(task)
	if err != nil {
		return err
	}
	err = n.rc.RPush(n.ctx, key, taskData).Err()
	if err != nil {
		return err
	}
	return nil
}

func (n KvRedis) Pop(key string) string {
	val, err := n.rc.LPop(n.ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
func (n KvRedis) SetEx(key string, value string, duration time.Duration) {
	err := n.rc.Set(n.ctx, key, value, duration).Err()
	if err != nil {
		return
	}
}
func (n KvRedis) Get(key string) string {
	val, err := n.rc.Get(n.ctx, key).Result()
	if err != nil {
		return ""
		/*panic(err)*/
	}
	return val
}
func (n KvRedis) LLen(key string) int64 {
	val, err := n.rc.LLen(n.ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
