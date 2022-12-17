package handler

import (
	"deps_per_service/internal/handler/mocks"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"testing"
)

func TestNewUserHandler(t *testing.T) {
	type args struct {
		userService UserService
	}

	tests := []struct {
		name string
		args args
		want *UserHandler
	}{
		{
			name: "TestNewUserHandler",
			args: args{
				userService: &mocks.MockUserService{},
			},
			want: &UserHandler{
				userService: &mocks.MockUserService{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHandler(tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHandler_CreateUser(t *testing.T) {
	type fields struct {
		userService UserService
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				userService: tt.fields.userService,
			}
			if err := h.CreateUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	type fields struct {
		userService UserService
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				userService: tt.fields.userService,
			}
			if err := h.DeleteUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	type fields struct {
		userService UserService
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				userService: tt.fields.userService,
			}
			if err := h.GetUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserHandler_GetUsers(t *testing.T) {
	type fields struct {
		userService UserService
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				userService: tt.fields.userService,
			}
			if err := h.GetUsers(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	type fields struct {
		userService UserService
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UserHandler{
				userService: tt.fields.userService,
			}
			if err := h.UpdateUser(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
