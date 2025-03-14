package service

import (
	"context"
	"github.com/JulyInSummer/cinematic/internal/app/service/domain"
	"github.com/JulyInSummer/cinematic/internal/app/storage"
	"go.uber.org/zap"
)

type cinematic struct {
	logger *zap.Logger
	movies storage.MoviesI
}

func NewCinematicService(logger *zap.Logger, movies storage.MoviesI) ServiceI {
	return &cinematic{
		logger: logger,
		movies: movies,
	}
}

func (c *cinematic) GetAll(ctx context.Context) ([]domain.Movie, error) {
	var movies []domain.Movie

	result, err := c.movies.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, movie := range result {
		movies = append(movies, domain.Movie{
			ID:       movie.ID,
			Title:    movie.Title,
			Director: movie.Director,
			Year:     movie.Year,
			Plot:     movie.Plot,
		})
	}

	return movies, nil
}

func (c *cinematic) GetByID(ctx context.Context, id int) (*domain.Movie, error) {
	result, err := c.movies.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Movie{
		ID:       result.ID,
		Title:    result.Title,
		Director: result.Director,
		Year:     result.Year,
		Plot:     result.Plot,
	}, nil
}

func (c *cinematic) Update(ctx context.Context, movie domain.Movie) (*domain.Movie, error) {
	result, err := c.movies.Update(ctx, movie.ToModel())
	if err != nil {
		return nil, err
	}

	return &domain.Movie{
		ID:       result.ID,
		Title:    result.Title,
		Director: result.Director,
		Year:     result.Year,
		Plot:     result.Plot,
	}, nil
}

func (c *cinematic) Delete(ctx context.Context, id int) error {
	return c.movies.Delete(ctx, id)
}
