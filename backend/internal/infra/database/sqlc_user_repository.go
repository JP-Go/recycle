package database

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/domain/valueobjects"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	queries *Queries
}

func NewUserRepository(queries *Queries) *UserRepository {
	return &UserRepository{
		queries: queries,
	}
}

func (ur *UserRepository) FindById(ctx context.Context, uid entities.UserId) (*entities.ValidatedUser, error) {
	u, err := ur.queries.FindUserById(ctx, uuid.UUID(uid))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repositories.ErrUserNotFound
		} else {
			return nil, err

		}
	}
	user := fromSqlcUserToUser(u)
	return user.Validate(), nil
}
func (ur *UserRepository) FindByEmail(ctx context.Context, email valueobjects.Email) (*entities.ValidatedUser, error) {
	u, err := ur.queries.FindUserByEmail(ctx, email.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repositories.ErrUserNotFound
		} else {
			return nil, err
		}

	}
	user := fromSqlcUserToUser(u)
	return user.Validate(), nil
}
func (ur *UserRepository) Create(ctx context.Context, user *entities.ValidatedUser) (*entities.ValidatedUser, error) {
	dbUser := fromUserToSqlcUser(*user.User)
	u, err := ur.queries.CreateUser(ctx, CreateUserParams{
		ID:                dbUser.ID,
		Email:             dbUser.Email,
		HashedPassword:    dbUser.HashedPassword,
		ConfirmationToken: dbUser.ConfirmationToken,
		Role:              dbUser.Role,
		CreatedBy:         dbUser.CreatedBy,
		UpdatedBy:         dbUser.UpdatedBy,
	})
	if err != nil {
		return &entities.ValidatedUser{}, err
	}
	saved := fromSqlcUserToUser(u)
	return saved.Validate(), nil
}

func (ur *UserRepository) IsEmailUnique(ctx context.Context, email valueobjects.Email) (bool, error) {
	user, err := ur.queries.FindUserByEmail(ctx, email.String())
	if err == nil {
		return user.DeletedAt.Valid, nil
	}
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return true, nil
	}
	return false, err
}

func fromSqlcUserToUser(u RecycleUser) entities.User {
	e, _ := valueobjects.NewEmail(u.Email)
	return entities.User{
		ID:                entities.UserId(u.ID),
		Email:             e,
		HashedPassword:    u.HashedPassword,
		ConfirmedAt:       fromTimestampzToTime(u.ConfirmedAt),
		ConfirmationToken: u.ConfirmationToken.String,
		Role:              entities.UserRole(u.Role),
		CreatedAt:         u.CreatedAt.Time,
		UpdatedAt:         u.UpdatedAt.Time,
		DeletedAt:         fromTimestampzToTime(u.DeletedAt),
		CreatedBy:         entities.UserId(u.CreatedBy),
		UpdatedBy:         entities.UserId(u.UpdatedBy),
	}
}

func fromUserToSqlcUser(u entities.User) RecycleUser {
	return RecycleUser{
		ID:                uuid.UUID(u.ID),
		Email:             u.Email.String(),
		HashedPassword:    u.HashedPassword,
		ConfirmedAt:       fromTimeToTimestampz(u.ConfirmedAt),
		ConfirmationToken: pgtype.Text{String: u.ConfirmationToken, Valid: true},
		Role:              string(u.Role),
		CreatedAt:         fromTimeToTimestampz(&u.CreatedAt),
		UpdatedAt:         fromTimeToTimestampz(&u.UpdatedAt),
		DeletedAt:         pgtype.Timestamptz{},
		CreatedBy:         uuid.UUID(u.CreatedBy),
		UpdatedBy:         uuid.UUID(u.UpdatedBy),
	}
}
