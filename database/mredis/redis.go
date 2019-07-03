package mredis

import (
	"time"

	"github.com/garyburd/redigo/redis"
	// "github.com/chasex/redis-go-cluster"
)

// RedisClient Client
var RedisClient *redis.Pool

// NewRedis 初始化Redis
func NewRedis(redisHost, Password string, redisDB int, network string) {

	client := &redis.Pool{
		MaxIdle:     50,
		MaxActive:   0,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, redisHost)
			if err != nil {
				return nil, err
			}
			c.Do("AUTH", Password)
			// 选择db
			c.Do("SELECT", redisDB)
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	RedisClient = client
}
