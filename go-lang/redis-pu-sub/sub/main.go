package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"fmt"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "redis-19958.c84.us-east-1-2.ec2.cloud.redislabs.com:19958",
	Username: "default",
	Password: "rvekWSq6H767ZfKf6cBAaUTDCLfflBvJ",
})


type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}


func main() {

	err := godotenv.Load() //
	if err != nil {	
		log.Fatal(err)
	}

	subscriber := redisClient.Subscribe(ctx, "send-user-data")
    user := User{}

    for {
        msg, err := subscriber.ReceiveMessage(ctx)
        if err != nil {
            panic(err)
        }

        if err := json.Unmarshal([]byte(msg.Payload), &user);
        err != nil{
            panic(err)
        }

        fmt.Println("Received message from " + msg.Channel + " channel.")
        fmt.Println("%+v\n", user)
    }

}
