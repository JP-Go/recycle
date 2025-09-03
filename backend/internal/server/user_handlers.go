package server

import (
	"backend/internal/application"
	database "backend/internal/database/postgresql"
	"backend/internal/services/hasher"
	"crypto/sha256"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type handleCreateUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type handleCreateUserOutput struct {
	CreatedAt time.Time `json:"created_at"`
}

func (s *Server) handleCreateUser(c echo.Context) error {
	in := new(handleCreateUserInput)
	validate := validator.New()
	if err := c.Bind(&in); err != nil {
		he := NewHttpError(http.StatusBadRequest, "Invalid payload", ErrInvalidPayload)
		return c.JSON(he.StatusCode, he)
	}

	if err := validate.Struct(in); err != nil {
		var errs validator.ValidationErrors
		var he httpError
		if errors.As(err, &errs) {
			verrs := []struct{ Field, Message string }{}
			for _, ve := range errs {
				verrs = append(verrs, struct{ Field, Message string }{
					Field:   ve.Field(),
					Message: ve.Error(),
				})
			}
			he = NewHttpError(http.StatusBadRequest, verrs, ErrInvalidPayload)
		} else {
			he = NewHttpError(http.StatusBadRequest, err.Error(), ErrInvalidPayload)
		}
		return c.JSON(he.StatusCode, he)
	}

	hasher := hasher.NewBcryptHasher(hasher.WithCost(14))
	hashedPassword, err := hasher.Compute(in.Password)
	if err != nil {
		he := NewHttpError(http.StatusInternalServerError, "Error calculating hash", ErrInternal)
		return c.JSON(he.StatusCode, he)
	}
	userId := uuid.New()
	// TODO: generate tokens with expiration
	sha := sha256.New()
	confirmationToken := sha.Sum([]byte(time.Now().String()))

	// TODO: send user confirmation token via email
	_, err = s.queries.CreateUser(c.Request().Context(), database.CreateUserParams{
		ID:                userId,
		Email:             in.Email,
		HashedPassword:    hashedPassword,
		ConfirmationToken: pgtype.Text{String: string(confirmationToken), Valid: true},
		Role:              string(application.RoleUser),
		CreatedBy:         userId,
		UpdatedBy:         userId,
	})
	if err != nil {
		log.Println(err)
		he := NewHttpError(http.StatusInternalServerError, "Error saving data", ErrInternal)
		return c.JSON(he.StatusCode, he)
	}
	return c.JSON(http.StatusCreated, handleCreateUserOutput{
		CreatedAt: time.Now(),
	})
}
