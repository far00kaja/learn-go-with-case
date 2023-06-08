package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(redisUrl string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		log.Fatal(err.Error())
	}

	client := redis.NewClient(opt)

	return client, nil
}

func RedisConnectInit() (*redis.Client, string, error) {
	REDIS_URL := os.Getenv("REDIS_URL")

	if REDIS_URL == "" {
		return nil, "", errors.New("missing REDIS_URL from env")
	}

	fmt.Println(REDIS_URL)
	redisConnect, err := NewRedisClient(REDIS_URL)
	if err != nil {
		return redisConnect, "", errors.New("cannot connected to Redis")
	}

	return redisConnect, "Connected to redis successfully", nil

}

func RedisConnect() *redis.Client {

	redisConnect, _, err := RedisConnectInit()
	if err != nil {
		panic(err)
	}

	return redisConnect
}
