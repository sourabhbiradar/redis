package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

type Product struct {
	ID   string
	Name string `json:"name"`
	Qty  string `json:"qty"`
}

var ctx = context.Background()

func main() {
	fmt.Println("---Redis---")

	// creating Redis Client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // default addr
		Password: "",               // no password bt reqd. at prodution level
		DB:       0,                // default DB 0
	})

	// checking client
	pong, err := client.Ping(ctx).Result()
	log.Println(pong, err)

	// initializing Product struct
	product := Product{
		ID:   os.Args[1], // pass arg in console
		Name: "chip",
		Qty:  "20",
	}

	// marshaling to JSON
	obj, err := json.Marshal(product)
	if err != nil {
		log.Printf("Error marshalling :%v\n", err)
	}

	// Set : updates or creates if non-existant on that ID
	// setting product to tht ID using redis.Set

	err = client.Set(ctx, product.ID, obj, 0).Err()
	if err != nil {
		log.Println(err)
	}

	// Get : gets data
	val, err := client.Get(ctx, product.ID).Result()
	if err != nil {
		log.Println(err)
	}
	log.Println("Product : ", val)

	// Del : delete
	res := client.Del(ctx, "344") // delete ID = 344
	log.Println(res)              // returns delete count

}
