package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	RedisKeyPrefix string = "redirection#"
)

type RedisConnection struct {
	*redis.Pool
}

func OpenConnection(server string) *RedisConnection {
	return &RedisConnection{newPool(server)}
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,

		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (self *RedisConnection) Exist(fqdn string) bool {
	conn := self.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", generateRedisKey(fqdn)))
	HandleErr(err)

	return exists
}

func (self *RedisConnection) Create(fqdn, target string) {
	conn := self.Get()
	defer conn.Close()

	key := generateRedisKey(fqdn)

	_, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(Redirection{
		From: fqdn, To: target, Count: 0,
	})...)
	HandleErr(err)
}

func (self *RedisConnection) GetAndIncrementCount(fqdn string) *Redirection {
	conn := self.Get()
	defer conn.Close()

	redirection := Redirection{}
	key := generateRedisKey(fqdn)

	data, err := redis.Values(conn.Do("HGETALL", key))
	HandleErr(err)
	HandleErr(redis.ScanStruct(data, &redirection))

	_, err = conn.Do("HINCRBY", key, "count", 1)
	HandleErr(err)

	return &redirection
}

func generateRedisKey(fqdn string) string {
	return RedisKeyPrefix + fqdn
}

type Redirection struct {
	From  string `redis:"from"`
	To    string `redis:"to"`
	Count int    `redis:"count"`
}
