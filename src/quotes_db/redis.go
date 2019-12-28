package quotes_db

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"os"
	"reflect"
	"strconv"
)

func RedisConnect() *redis.Client {

	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		return nil
	}

	var client = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       redisDB,
	})

	return client
}

func QuotesMapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()

	return keys[rand.Intn(len(keys))].Interface()
}

func GetQuotes(hashSetName string) map[string]string {

	client := RedisConnect()

	val, err := client.HGetAll(hashSetName).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Got quote", val)

	client.Close()

	return val
}