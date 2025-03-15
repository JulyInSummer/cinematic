package postgres

import (
	"context"
	"github.com/JulyInSummer/cinematic/internal/app/storage"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User struct {
	logger *zap.Logger
	db     *gorm.DB
	config *Config
}

func NewUserStorage(logger *zap.Logger, db *gorm.DB, config *Config) storage.UsersI {
	return &User{
		logger: logger,
		db:     db,
		config: config,
	}
}

func (u *User) Create(ctx context.Context, user models.User) error {
	method := "User.Create"

	result := u.db.WithContext(ctx).Save(&user)
	if err := result.Error; err != nil {
		u.logger.Error(method, zap.Error(err))
		return err
	}

	return nil
}

func (u *User) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	method := "User.GetByEmail"
	var user models.User

	result := u.db.WithContext(ctx).First(&models.User{}, "email = ?", email).Scan(&user)
	if err := result.Error; err != nil {
		u.logger.Error(method, zap.Error(err))
		return nil, err
	}

	return &user, nil
}
