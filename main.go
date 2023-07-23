package main

import (
	"flag"

	"github.com/1eedaegon/fast-booking-svc-practice/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.String("port", ":8888", "Port of ")
	flag.Parse()

	app := fiber.New()
	appV1 := app.Group("/app/v1")
	appV1.Get("/user", api.HandleGetUsers)
	appV1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*port)

}
