package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/", handleRequest)

	port := os.Getenv("PORT")

	fmt.Printf("Server is running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	count, err := rdb.Incr(ctx, "access_count").Result()
	if err != nil {
		http.Error(w, "Error incrementing access count", http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Access Count: %d", count)
	fmt.Fprint(w, message)
}
