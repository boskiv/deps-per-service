package repository

import (
	"deps_per_service/internal/core"
	"deps_per_service/internal/service"
	"gorm.io/gorm"
)

type userRepo struct {
	BaseRepo
}

func (u *userRepo) Get(id int) (*core.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) GetByEmail(email string) (*core.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Create(user *core.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Update(user *core.UserServiceDTO) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(db *gorm.DB) service.UserRepo {
	return &userRepo{
		BaseRepo: BaseRepo{
			db: db,
		},
	}
}
