package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ok, err := client.Set("wilson", "123", time.Hour).Result()
	if err != nil {
		log.Fatal("Failed to set: ", err)
	}

	log.Println("Wilson is setted: ", ok)

	result, err := client.Get("wilson").Result()
	if err != nil {
		log.Fatal("Failed to get: ", err)
	}
	fmt.Println("O valor de wilson é ", result)
}
