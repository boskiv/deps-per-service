package service

import (
	"deps_per_service/internal/core"
)

type ProfileRepo interface {
	GetProfile(userID int) (*core.Profile, error)
	UpdateProfile(userID int, profile *core.Profile) error
	CreateProfile(userID int, profile *core.Profile) error
}

type ProfileService struct {
	profileRepo ProfileRepo
}

func (p *ProfileService) GetProfile(userID int) (*core.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProfileService) CreateProfile(profile *core.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProfileService) UpdateProfile(profile *core.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProfileService) DeleteProfile(userID int) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileService(profileRepo ProfileRepo) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
	}
}
