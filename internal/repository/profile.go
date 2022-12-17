package repository

import (
	"deps_per_service/internal/core"
	"deps_per_service/internal/service"
	"gorm.io/gorm"
)

type profileRepo struct {
	BaseRepo
}

func (p *profileRepo) GetProfile(userID int) (*core.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p *profileRepo) UpdateProfile(userID int, profile *core.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p *profileRepo) CreateProfile(userID int, profile *core.Profile) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileRepo(db *gorm.DB) service.ProfileRepo {
	return &profileRepo{
		BaseRepo: BaseRepo{
			db: db,
		},
	}
}
