package app

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
	"os"
)
var ctx = context.Background()

func greet(writer http.ResponseWriter, request *http.Request) {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	_, _ = fmt.Fprintf(writer, pong)
	client.HSet(ctx, "", "")
}
