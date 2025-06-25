package db

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

/*
 * @brief InitRedis initializes a Redis client for caching URL lookups.
 *
 * This function creates a Redis client connected to the Redis server at
 * "redis:6379". The client is stored in the package variable RDB.
 */
func InitRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr: "redis:6379",
    })
}