package handler

import (
	"deps_per_service/internal/core"
	"github.com/gofiber/fiber/v2"
)

type ProfileService interface {
	GetProfile(userID int) (*core.Profile, error)
	CreateProfile(profile *core.Profile) error
	UpdateProfile(profile *core.Profile) error
	DeleteProfile(userID int) error
}

type ProfileHandler struct {
	profileService ProfileService
}

func NewProfileHandler(profileService ProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

func (h *ProfileHandler) GetProfile(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *ProfileHandler) CreateProfile(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *ProfileHandler) UpdateProfile(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}

func (h *ProfileHandler) DeleteProfile(ctx *fiber.Ctx) error {
	// TODO implement me
	panic("implement me")
}
