package service

import (
	"context"
	"github.com/JulyInSummer/cinematic/internal/app/pkg/jwt"
	"github.com/JulyInSummer/cinematic/internal/app/service/domain"
	"github.com/JulyInSummer/cinematic/internal/app/storage"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type cinematic struct {
	logger *zap.Logger
	conf   *Config
	movies storage.MoviesI
	users  storage.UsersI
}

func NewCinematicService(logger *zap.Logger, movies storage.MoviesI, users storage.UsersI, conf *Config) ServiceI {
	return &cinematic{
		logger: logger,
		movies: movies,
		users:  users,
		conf:   conf,
	}
}

func (c *cinematic) CreateUser(ctx context.Context, user domain.User) error {
	method := "cinematic.CreateUser"
	var userData domain.User

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		c.logger.Error(method, zap.Error(err))
		return err
	}

	userData.Email = user.Email
	userData.Password = string(password)

	if err = c.users.Create(ctx, userData.ToModel()); err != nil {
		c.logger.Error(method, zap.Error(err))
		return err
	}

	return nil
}

func (c *cinematic) GetUserTokenByEmail(ctx context.Context, credentials domain.User) (*domain.Token, error) {
	method := "cinematic.GetUserByEmail"

	var user domain.User

	res, err := c.users.GetByEmail(ctx, credentials.Email)
	if err != nil {
		c.logger.Error(method, zap.Error(err))
		return nil, err
	}

	user.Email = res.Email
	user.Password = res.Password

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.logger.Error(method, zap.Error(err))
		return nil, err
	}

	token, err := jwt.CreateToken(user.Email, c.conf.Secret)
	if err != nil {
		c.logger.Error(method, zap.Error(err))
		return nil, err
	}

	return &domain.Token{Token: "Bearer " + token}, nil
}

func (c *cinematic) Create(ctx context.Context, movie domain.Movie) error {
	return c.movies.Create(ctx, movie.ToModel())
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

func (c *cinematic) Update(ctx context.Context, movie domain.Movie) error {
	return c.movies.Update(ctx, movie.ToModel())

}

func (c *cinematic) Delete(ctx context.Context, id int) error {
	return c.movies.Delete(ctx, id)
}
