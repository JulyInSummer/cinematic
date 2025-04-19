package postgres

import (
	"context"
	"fmt"

	"github.com/JulyInSummer/cinematic/internal/app/storage"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/query"
	"github.com/jackc/pgx/v5"

	//_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Storage struct {
	logger *zap.Logger
	db     *pgx.Conn
	config *Config
}

func getConnString(conf *Config) string {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresUser,
		conf.PostgresDatabase,
		conf.PostgresPassword,
	)

	return connStr
}

func NewStorage(logger *zap.Logger, db *pgx.Conn, conf *Config) (storage.MoviesI, error) {

	return &Storage{
		logger: logger,
		db:     db,
		config: conf,
	}, nil
}

func (s *Storage) Create(ctx context.Context, movie models.Movie) error {
	method := "Storage.Create"

	_, err := s.db.Exec(ctx, query.CreateMovie, movie.Title, movie.Director, movie.Year, movie.Plot)
	if err != nil {
		s.logger.Error(method, zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) GetAll(ctx context.Context) ([]models.Movie, error) {
	method := "Storage.GetAll"

	var movies []models.Movie
	rows, err := s.db.Query(ctx, query.GetAllMovies)
	if err != nil {
		s.logger.Error(method, zap.Error(err))
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var movie models.Movie
		if err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Director,
			&movie.Year,
			&movie.Plot,
		); err != nil {
			s.logger.Error(method, zap.Error(err))
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (s *Storage) GetByID(ctx context.Context, id uint64) (*models.Movie, error) {
	method := "Storage.GetByID"

	var movie models.Movie
	if err := s.db.QueryRow(ctx, query.GetMovieByID, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Director,
		&movie.Year,
		&movie.Plot,
	); err != nil {
		s.logger.Error(method, zap.Error(err))
		return nil, err
	}

	return &movie, nil
}

func (s *Storage) Update(ctx context.Context, movie models.Movie) error {
	method := "Storage.Update"

	_, err := s.GetByID(ctx, movie.ID)
	if err != nil {
		s.logger.Error(method, zap.Error(err))
		return err
	}

	params := pgx.NamedArgs{
		"id":       movie.ID,
		"title":    movie.Title,
		"director": movie.Director,
		"year":     movie.Year,
		"plot":     movie.Plot,
	}

	_, err = s.db.Exec(ctx, query.UpdateMovie, params)
	if err != nil {
		s.logger.Error(method, zap.Error(err))
		return err
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, id uint64) error {
	method := "Storage.Delete"

	_, err := s.GetByID(ctx, id)
	if err != nil {
		s.logger.Error(method, zap.Error(err))
		return err
	}

	_, err = s.db.Exec(ctx, query.DeleteMovie, id)
	if err != nil {
		s.logger.Error(method, zap.Error(err))
		return err
	}

	return nil
}
