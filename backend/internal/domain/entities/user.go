package entities

import (
	"backend/internal/domain/valueobjects"
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserRole string

const RoleUser UserRole = "user"
const RoleCompanyManager UserRole = "company_manager"
const RoleAdmin UserRole = "admin"

type UserId uuid.UUID

type User struct {
	ID                UserId
	Email             valueobjects.Email
	HashedPassword    string
	ConfirmedAt       *time.Time
	ConfirmationToken string
	Role              UserRole
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
	CreatedBy         UserId
	UpdatedBy         UserId
}

type ValidatedUser struct {
	*User
	valid bool
}

var ErrEmailNotUnique = errors.New("Email not unique")

func SelfCreateNewUser(emailStr, confirmationToken, hashedPassword string) (*ValidatedUser, error) {
	email, err := valueobjects.NewEmail(emailStr)
	if err != nil {
		return &ValidatedUser{valid: false}, err
	}
	id := UserId(uuid.New())

	return &ValidatedUser{
		User: &User{
			ID:                id,
			Email:             email,
			HashedPassword:    hashedPassword,
			ConfirmedAt:       nil,
			ConfirmationToken: confirmationToken,
			Role:              RoleUser,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			DeletedAt:         nil,
			CreatedBy:         id,
			UpdatedBy:         id,
		},
	}, nil
}

func NewUser(emailStr, confirmationToken, hashedPassword string, createdBy uuid.UUID) (*ValidatedUser, error) {
	email, err := valueobjects.NewEmail(emailStr)
	if err != nil {
		return &ValidatedUser{valid: false}, err
	}

	return &ValidatedUser{
		User: &User{
			ID:                UserId(uuid.New()),
			Email:             email,
			HashedPassword:    hashedPassword,
			ConfirmedAt:       nil,
			ConfirmationToken: confirmationToken,
			Role:              RoleUser,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			DeletedAt:         nil,
			CreatedBy:         UserId(createdBy),
			UpdatedBy:         UserId(createdBy),
		},
	}, nil
}

func NewCompanyOwner(emailStr, confirmationToken, hashedPassword string, createdBy uuid.UUID) (*ValidatedUser, error) {
	email, err := valueobjects.NewEmail(emailStr)
	if err != nil {
		return &ValidatedUser{valid: false}, nil
	}

	return &ValidatedUser{
		User: &User{
			ID:                UserId(uuid.New()),
			Email:             email,
			HashedPassword:    hashedPassword,
			ConfirmedAt:       nil,
			ConfirmationToken: confirmationToken,
			Role:              RoleCompanyManager,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			DeletedAt:         nil,
			CreatedBy:         UserId(createdBy),
			UpdatedBy:         UserId(createdBy),
		},
	}, nil
}

func NewAdmin(emailStr, confirmationToken, hashedPassword string, createdBy uuid.UUID) (*ValidatedUser, error) {
	email, err := valueobjects.NewEmail(emailStr)
	if err != nil {
		return &ValidatedUser{valid: false}, nil
	}

	return &ValidatedUser{
		valid: true,
		User: &User{
			ID:                UserId(uuid.New()),
			Email:             email,
			HashedPassword:    hashedPassword,
			ConfirmedAt:       nil,
			ConfirmationToken: confirmationToken,
			Role:              RoleCompanyManager,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			DeletedAt:         nil,
			CreatedBy:         UserId(createdBy),
			UpdatedBy:         UserId(createdBy),
		},
	}, nil
}
func (u *User) Validate() *ValidatedUser {
	if !u.Email.IsValid() {
		return &ValidatedUser{u, false}
	}
	return &ValidatedUser{u, true}
}

func (vu ValidatedUser) IsValid() bool {
	return vu.valid
}
