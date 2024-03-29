package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func connectDB() *redis.Client {

	redisAddr := "redis-service.socialhub.svc.cluster.local:6379"
	if r := os.Getenv("REDIS_ADDR"); r != "" {
		redisAddr = r
	}

	redisPass := ""
	if r := os.Getenv("REDIS_PASS"); r != "" {
		redisAddr = r
	}

	// defines redis connection
	rdb = redis.NewClient(&redis.Options{
		// Addr: "localhost:6379",
		// Addr:     "redis:6379",
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
	})

	// simple ping / connection check
	pong, err := rdb.Ping().Result()
	log.Println(pong, err)
	return rdb
}

func sendToRedis(user User) {
	var m = make(map[string]interface{})
	m["username"] = user.Username
	m["password"] = user.Password
	rdb.HMSet(user.Id, m)

}
