package main

import (
	"log"
	"time"

	"gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ok, err := client.Set("wilson", "valor", time.Hour).Result()
	if err != nil {
		log.Fatal("Failed to set: ", err)
	}

	log.Println("Wilson is setted: ", ok)
}
