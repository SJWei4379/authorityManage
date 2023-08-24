package global

import "github.com/gomodule/redigo/redis"

// 创建redis连接池
var (
	RedisPool *redis.Pool
)
