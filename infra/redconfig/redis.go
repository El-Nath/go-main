package redconfig

import (
	"github.com/go-redis/redis"
)

type Env struct {
	RC *redis.Client
}

func ConnectRedis(a string, p string, dB int) (*redis.Client, error) {

	c := redis.NewClient(&redis.Options{
		Addr:     a,
		Password: p,
		DB:       dB,
	})

	err := c.Ping().Err()
	if err != nil {
		return c, err

	}
	return c, nil
}
