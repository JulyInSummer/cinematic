package storage

import (
	"context"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"
)

type MoviesI interface {
	Create(ctx context.Context, movie models.Movie) error
	GetAll(ctx context.Context) ([]models.Movie, error)
	GetByID(ctx context.Context, id uint64) (*models.Movie, error)
	Update(ctx context.Context, movie models.Movie) error
	Delete(ctx context.Context, id uint64) error
}

type UsersI interface {
	Create(ctx context.Context, user models.User) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}
