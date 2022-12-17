package service

import (
	"deps_per_service/internal/core"
	"deps_per_service/internal/handler"
)

type ProfileRepo interface {
	GetProfile(userID int) (*core.Profile, error)
	UpdateProfile(userID int, profile *core.Profile) error
	CreateProfile(userID int, profile *core.Profile) error
}

type profileService struct {
	profileRepo ProfileRepo
}

func (p *profileService) GetProfile(userID int) (*core.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p *profileService) CreateProfile(profile *core.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p *profileService) UpdateProfile(profile *core.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p *profileService) DeleteProfile(userID int) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileService(profileRepo ProfileRepo) handler.ProfileService {
	return &profileService{
		profileRepo: profileRepo,
	}
}
