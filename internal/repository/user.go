package repository

import (
	"deps_per_service/internal/core"
	"gorm.io/gorm"
)

type UserRepo struct {
	BaseRepo
}

func (u *UserRepo) Get(id int) (*core.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) GetByEmail(email string) (*core.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Create(user *core.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Update(user *core.UserServiceDTO) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		BaseRepo: BaseRepo{
			db: db,
		},
	}
}
