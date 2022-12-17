package core

type User struct {
	ID       int
	Username string
	Password string
}

// UserServiceDTO is a DTO for UserService. It is used to transfer data between the service and the handler.
type UserServiceDTO struct {
	ID       int
	Username string `json:"username" validate:"required" mapstructure:"username,omitempty"`
	Password string `json:"-" validate:"required" mapstructure:"password,omitempty"`
}
