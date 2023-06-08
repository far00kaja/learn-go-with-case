package lib

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/far00kaja/learn-go-with-case/config"
	"github.com/go-redis/redis/v8"
)

type Service interface {
	SetRedisNew(key string, value interface{}) (*redis.StatusCmd, error)
	GetRedisFromKey(key string) (int, error)
}

func SetRedis(key string, value interface{}) (*redis.StatusCmd, error) {
	// ttl := time.Duration(1) * time.Hour

	setRedis := config.RedisConnect().Set(context.Background(), key, value, 0)
	if err := setRedis.Err(); err != nil {
		return setRedis, err
		// return setRedis, errors.New("unable to SET data in redis")
	}

	successMsg := fmt.Sprintf("Successfully save new key to redis, key :%s", key)

	log.Println(successMsg)

	return setRedis, nil
}

func GetRedisFromKeyIntValue(key string) (int, error) {
	getRedis := config.RedisConnect().Get(context.Background(), key)
	if err := getRedis.Err(); err != nil {
		return 0, errors.New("unable to GET data in redis")
	}

	data, err := getRedis.Int()
	if err != nil {
		return 0, errors.New("unable to GET data in redis")
	}

	return data, nil
}

func GetRedisFromKeyStrValue(key string) (string, error) {
	getRedis := config.RedisConnect().Get(context.Background(), key)
	if err := getRedis.Err(); err != nil {
		return "", errors.New("unable to GET data in redis")
	}

	data := getRedis.String()

	return data, nil
}

func DeleteRedisFromKey(key string) (*redis.IntCmd, error) {
	deleteRedis := config.RedisConnect().Del(context.Background(), key)
	// deleteRedis := config.RedisConnect().Del(RedisConnect().Context(), key).Result()
	if err := deleteRedis.Err(); err != nil {
		// logger.Printf(deleteRedis.Err().Error())
		return deleteRedis, errors.New("unable to Delete data from redis")
	}

	return deleteRedis, nil
}
