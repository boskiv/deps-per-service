package handler

import "deps_per_service/internal/core"

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
