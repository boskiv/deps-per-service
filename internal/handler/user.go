package handler

import (
	"deps_per_service/internal/core"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	// CreateUser creates a new user.
	CreateUser(user *core.User) error
	// GetUser returns the user data for the given user ID.
	GetUser(userID int) (*core.User, error)
	// UpdateUser updates the user data for the given user ID.
	UpdateUser(userID int, user *core.User) error
	// DeleteUser deletes the user data for the given user ID.
	DeleteUser(userID int) error
}

type UserHandler struct {
	userService UserService
}

func (h *UserHandler) GetUsers(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
