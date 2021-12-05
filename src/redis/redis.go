package redis

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatalln("Redis error", err)
	}

	log.Println("Redis PING -> " + pong)
}

func GetInstance() *redis.Client {
	if client == nil {
		Init()
	}
	return client
}
