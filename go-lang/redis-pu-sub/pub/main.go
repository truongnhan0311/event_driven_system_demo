package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
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

	err := godotenv.Load() // 👈 load .env file
	if err != nil {	
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			panic(err)
		}

		payload, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}


		if err := redisClient.Publish(ctx, "send-user-data", payload).Err(); err != nil {
			panic(err)
		}

		return c.SendStatus(200)


	})

	app.Listen(":3000")
}
