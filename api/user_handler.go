package api

import (
	"github.com/1eedaegon/fast-booking-svc-practice/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	// u := types.User{ID: "1", FirstName: "lee", LastName: "daeon"}
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	// u := types.User{ID: "2", FirstName: "An", LastName: "sunghyun"}
	return c.JSON(u)
}