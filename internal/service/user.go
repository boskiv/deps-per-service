package service

import (
	"deps_per_service/internal/core"
)

type UserRepo interface {
	// Get returns a user by ID.
	Get(id int) (*core.User, error)
	// GetByEmail returns a user by email.
	GetByEmail(email string) (*core.User, error)
	// Create creates a new user.
	Create(user *core.User) error
	// Update updates a user.
	Update(user *core.UserServiceDTO) error
	// Delete deletes a user.
	Delete(id int) error
}

type UserService struct {
	userRepo UserRepo
}

func (u *UserService) CreateUser(user *core.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) GetUser(userID int) (*core.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UpdateUser(userID int, user *core.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) DeleteUser(userID int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
