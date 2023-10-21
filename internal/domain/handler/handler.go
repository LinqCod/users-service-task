package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/linqcod/users-service-task/internal/domain/handler/dto"
	"github.com/linqcod/users-service-task/internal/domain/model"
	"github.com/linqcod/users-service-task/pkg/response"
	"go.uber.org/zap"
	"net/http"
)

type UserService interface {
	GetUsers() *model.UserList
	CreateUser(userDto dto.CreateUserRequest) (string, error)
	GetUser(id string) (model.User, error)
	UpdateUser(id string, user dto.UpdateUserRequest) error
	DeleteUser(id string) error
}

type UserHandler struct {
	logger  *zap.SugaredLogger
	service UserService
}

func NewUserHandler(logger *zap.SugaredLogger, service UserService) *UserHandler {
	return &UserHandler{
		logger:  logger,
		service: service,
	}
}

func (h UserHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetUsers()
	render.JSON(w, r, users)
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := dto.CreateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	id, err := h.service.CreateUser(request)
	if err != nil {
		h.logger.Errorf("error while creating user: %v", err)
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.service.GetUser(id)
	if err != nil {
		h.logger.Errorf("error while getting user: %v", err)
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, user)
}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	request := dto.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	if err := h.service.UpdateUser(id, request); err != nil {
		h.logger.Errorf("error while updating user: %v", err)
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.service.DeleteUser(id); err != nil {
		h.logger.Errorf("error while getting user: %v", err)
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
