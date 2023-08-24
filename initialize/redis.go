package initialize

import (
	"authoritymanage/global"
	"github.com/gomodule/redigo/redis"
	"log"
)

func RedisInit() {
	if global.RedisPool == nil {
		global.RedisPool = &redis.Pool{ //实例化一个连接池
			MaxIdle:     3, //最初的连接数量
			MaxActive:   0, //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
			IdleTimeout: 0, //连接关闭时间
			Dial: func() (redis.Conn, error) { //要连接的redis数据库
				dial, err := redis.Dial("tcp", "101.42.222.26:6379", redis.DialPassword("wsj@admin"))
				if err != nil {
					log.Fatalln("redis连接失败", err.Error())
				}
				return dial, err
			},
		}
		global.RedisPool.Get()
	}
}
