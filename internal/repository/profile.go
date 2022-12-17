package repository

import (
	"deps_per_service/internal/core"
	"gorm.io/gorm"
)

type ProfileRepo struct {
	BaseRepo
}

func (p *ProfileRepo) GetProfile(userID int) (*core.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProfileRepo) UpdateProfile(userID int, profile *core.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProfileRepo) CreateProfile(userID int, profile *core.Profile) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileRepo(db *gorm.DB) *ProfileRepo {
	return &ProfileRepo{
		BaseRepo: BaseRepo{
			db: db,
		},
	}
}
