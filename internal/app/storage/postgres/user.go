package postgres

import (
	"context"
	"github.com/JulyInSummer/cinematic/internal/app/storage"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/query"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type User struct {
	logger *zap.Logger
	db     *pgx.Conn
	config *Config
}

func NewUserStorage(logger *zap.Logger, db *pgx.Conn, config *Config) storage.UsersI {
	return &User{
		logger: logger,
		db:     db,
		config: config,
	}
}

func (u *User) Create(ctx context.Context, user models.User) error {
	method := "User.Create"

	_, err := u.db.Exec(ctx, query.CreateUser, user.Email, user.Password)
	if err != nil {
		u.logger.Error(method, zap.Error(err))
		return err
	}

	return nil
}

func (u *User) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	method := "User.GetByEmail"
	var user models.User

	if err := u.db.QueryRow(ctx, query.GetUserByEmail, email).Scan(&user.ID, &user.Email, &user.Password); err != nil {
		u.logger.Error(method, zap.Error(err))
		return nil, err
	}

	return &user, nil
}
