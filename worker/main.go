package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "redis_db:6379"})
	ctx := context.Background()
	subscriber := rdb.Subscribe(ctx, "number")

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			fmt.Printf("error while recieving message: %v\n", err)
			continue
		}

		var req struct {
			Value int `json:"value"`
		}

		if err := json.Unmarshal([]byte(msg.Payload), &req); err != nil {
			fmt.Printf("error unmarshaling message")
			continue
		}

		_, err = rdb.HSet(ctx, "numbers", map[string]interface{}{fmt.Sprintf("%v", req.Value): req.Value}).Result()
		if err != nil {
			fmt.Printf("error while inserting number")
			return
		}
	}
}
