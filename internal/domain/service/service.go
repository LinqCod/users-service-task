package service

import (
	"github.com/linqcod/users-service-task/internal/domain/handler/dto"
	"github.com/linqcod/users-service-task/internal/domain/model"
	"time"
)

type UserRepository interface {
	GetUsers() *model.UserList
	CreateUser(user model.User) (string, error)
	GetUser(id string) (model.User, error)
	UpdateUser(id string, user dto.UpdateUserRequest) error
	DeleteUser(id string) error
}

type UserService struct {
	repository UserRepository
}

func NewUserRepository(repo UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (s UserService) GetUsers() *model.UserList {
	return s.repository.GetUsers()
}

func (s UserService) CreateUser(userDto dto.CreateUserRequest) (string, error) {
	user := model.User{
		CreatedAt:   time.Now(),
		DisplayName: userDto.DisplayName,
		Email:       userDto.Email,
	}

	return s.repository.CreateUser(user)
}

func (s UserService) GetUser(id string) (model.User, error) {
	return s.repository.GetUser(id)
}

func (s UserService) UpdateUser(id string, user dto.UpdateUserRequest) error {
	return s.repository.UpdateUser(id, user)
}

func (s UserService) DeleteUser(id string) error {
	return s.repository.DeleteUser(id)
}
