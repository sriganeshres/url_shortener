package db

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

func InitRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr: "redis:6379",
    })
}