package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"backend/internal/application/services"
	"backend/internal/application/usecases"
	"backend/internal/infra/database"
)

type Server struct {
	port int

	queries           *database.Queries
	createUserUseCase *usecases.CreateUserUsecase
}

func NewServer(queries *database.Queries) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	ur := database.NewUserRepository(queries)
	NewServer := &Server{
		port:              port,
		queries:           queries,
		createUserUseCase: usecases.NewCreateUserUsecase(ur, services.NewBcryptHasher(), services.NewUniqueEmailVerifier(ur)),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
