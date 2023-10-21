package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/linqcod/users-service-task/internal/domain/handler"
	"github.com/linqcod/users-service-task/internal/domain/repository"
	"github.com/linqcod/users-service-task/internal/domain/service"
	"github.com/linqcod/users-service-task/pkg/database"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func InitRouter(logger *zap.SugaredLogger, store *database.UserStore) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// init handler, repo and service
	userRepo := repository.NewUserRepository(store)
	userService := service.NewUserRepository(userRepo)
	userHandler := handler.NewUserHandler(logger, userService)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", userHandler.SearchUsers)
				r.Post("/", userHandler.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", userHandler.GetUser)
					r.Patch("/", userHandler.UpdateUser)
					r.Delete("/", userHandler.DeleteUser)
				})
			})
		})
	})

	return r
}
