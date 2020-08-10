package main

import (
	"context"
	"delayed-queue/task"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func main() {

	redisClient, err := newRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	repo := task.NewRepository(redisClient)
	service := task.NewService(repo)

	router := newRouter(service)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Print("server started on 8080")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("err: %v", err)
	}

}

func newRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
