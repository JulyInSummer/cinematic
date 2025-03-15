package service

import (
	"context"
	"github.com/JulyInSummer/cinematic/internal/app/service/domain"
)

type ServiceI interface {
	Create(ctx context.Context, movie domain.Movie) error
	GetAll(ctx context.Context) ([]domain.Movie, error)
	GetByID(ctx context.Context, id int) (*domain.Movie, error)
	Update(ctx context.Context, movie domain.Movie) error
	Delete(ctx context.Context, id int) error
}
