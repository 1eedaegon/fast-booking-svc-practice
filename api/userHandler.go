package api

import (
	"github.com/1eedaegon/fast-booking-svc-practice/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	u := types.User{ID: "1", FirstName: "lee", LastName: "daeon"}
	return c.JSON(u)
}
func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{ID: "2", FirstName: "An", LastName: "sunghyun"}
	return c.JSON(u)
}
