package redis

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type redisClient struct {
	client *redis.Client
	redisClientConfig
}

type redisClientConfig struct {
	host     string
	port     string
	password string
}

var rdb *redisClient
var once sync.Once
var ctx context.Context

func New(host string, args ...options) *redisClient {
	rdb = &redisClient{
		redisClientConfig: redisClientConfig{
			host:     host,
			port:     "",
			password: "",
		},
	}

	for _, v := range args {
		v(rdb)
	}

	ctx = context.Background()
	return rdb
}

func (r *redisClient) With(args ...options) {
	for _, v := range args {
		v(rdb)
	}
}

func (r *redisClient) Init() {
	once.Do(func() {
		if r.host == "" {
			panic(errors.New("redis host not defined"))
		}

		if r.port == "" {
			r.port = "6379"
		}

		config := redis.Options{
			Addr: r.host + ":" + r.port,
			DB:   0,
		}

		if r.password != "" {
			config.Password = r.password
		}

		rdb.client = redis.NewClient(&config)
	})

	if _, err := rdb.client.Ping(ctx).Result(); err != nil {
		panic(errors.New("redis connect failed:" + err.Error()))
	}

	fmt.Println("redis connect succeeded")
}
