package storage

import (
	"context"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"
)

type MoviesI interface {
	Create(ctx context.Context, movie models.Movie) error
	GetAll(ctx context.Context) ([]models.Movie, error)
	GetByID(ctx context.Context, id int) (*models.Movie, error)
	Update(ctx context.Context, movie models.Movie) error
	Delete(ctx context.Context, id int) error
}
