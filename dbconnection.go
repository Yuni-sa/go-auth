package main

import (
	// "encoding/json"
	"log"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func connectDB() *redis.Client {
	// defines redis connection
	rdb = redis.NewClient(&redis.Options{
		// Addr: "localhost:6379",
		// Addr:     "redis:6379",
		Addr:     "redis-service.socialhub.svc.cluster.local:6379",
		Password: "",
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

// func getFromRedis(data map[string]string, c *fiber.Ctx) error {
// 	var user User
//
// 	user.Username = data["username"]
// 	user.Password = rdb.HGet(data["username"], "password").Val()
//
// 	if user.Password == "" {
// 		c.Status(fiber.StatusNotFound)
// 		return c.JSON(fiber.Map{
// 			"message": "user not found",
// 		})
// 	}
//
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
// 		c.Status(fiber.StatusBadRequest)
// 		return c.JSON(fiber.Map{
// 			"message": "incorrect password",
// 		})
// 	}
// 	return (c.JSON(user))
// }
