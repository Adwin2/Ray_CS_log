package main

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)


func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
	hashKey := "exampleHash"
	field1 := "field1"
	value1 := "value1"
	expiration := 40 * time.Second
	
	// Set a field in the hash
	err := rdb.HSet(ctx, hashKey, field1, value1).Err()
	if err != nil {
		log.Printf("HSet Error: %s", err)
		return
	}
	
	// Set expiration for the hash field
	err = rdb.HExpire(ctx, hashKey, expiration, field1).Err()
	if err != nil {
		log.Printf("HExpire Error: %s", err)
		return
	}
}