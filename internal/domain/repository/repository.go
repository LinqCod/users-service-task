package repository

import (
	"github.com/linqcod/users-service-task/internal/common/errorTypes"
	"github.com/linqcod/users-service-task/internal/domain/handler/dto"
	"github.com/linqcod/users-service-task/internal/domain/model"
	"github.com/linqcod/users-service-task/pkg/database"
	"strconv"
)

type UserRepository struct {
	store *database.UserStore
}

func NewUserRepository(store *database.UserStore) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (r UserRepository) GetUsers() *model.UserList {
	return &r.store.List
}

func (r UserRepository) CreateUser(user model.User) (string, error) {
	r.store.Increment++

	id := strconv.Itoa(r.store.Increment)
	r.store.List[id] = user

	if err := r.store.UpdateDB(); err != nil {
		return "", err
	}

	return id, nil
}

func (r UserRepository) GetUser(id string) (model.User, error) {
	if _, ok := r.store.List[id]; !ok {
		return model.User{}, errorTypes.ErrUserNotFound
	}

	return r.store.List[id], nil
}

func (r UserRepository) UpdateUser(id string, user dto.UpdateUserRequest) error {
	if _, ok := r.store.List[id]; !ok {
		return errorTypes.ErrUserNotFound
	}

	u := r.store.List[id]
	u.DisplayName = user.DisplayName
	r.store.List[id] = u

	if err := r.store.UpdateDB(); err != nil {
		return err
	}

	return nil
}

func (r UserRepository) DeleteUser(id string) error {
	if _, ok := r.store.List[id]; !ok {
		return errorTypes.ErrUserNotFound
	}

	delete(r.store.List, id)

	if err := r.store.UpdateDB(); err != nil {
		return err
	}

	return nil
}
