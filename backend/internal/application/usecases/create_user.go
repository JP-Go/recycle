package usecases

import (
	"backend/internal/application/services"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/domain/valueobjects"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateUserUsecase struct {
	ur            repositories.UserRepository
	emailVerifier services.UniqueEmailVerifier
	hasher        services.Hasher
}

func NewCreateUserUsecase(ur repositories.UserRepository, hasher services.Hasher, verifier services.UniqueEmailVerifier) *CreateUserUsecase {

	return &CreateUserUsecase{
		ur:            ur,
		hasher:        hasher,
		emailVerifier: verifier,
	}
}

// TODO: handle validation errors
type CreateUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

type CreateUserOutput struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func (cu *CreateUserUsecase) HandleCreateUser(e echo.Context) error {
	val := validator.New(validator.WithRequiredStructEnabled())
	input := new(CreateUserInput)
	if err := e.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err := val.Struct(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	emailToVerify, err := valueobjects.NewEmail(input.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	emailOk, err := cu.emailVerifier.Verify(e.Request().Context(), emailToVerify)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !emailOk {
		return echo.NewHTTPError(http.StatusConflict, "Email in use")
	}

	passwordHash, err := cu.hasher.Compute(input.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// TODO: time-aware confirmation token
	userToCreate, err := entities.SelfCreateNewUser(input.Email, "token", passwordHash)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	created, err := cu.ur.Create(e.Request().Context(), userToCreate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, CreateUserOutput{
		ID:        uuid.UUID(created.ID).String(),
		CreatedAt: created.CreatedAt,
	})
}
