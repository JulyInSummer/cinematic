package postgres

import (
	"context"
	"fmt"
	"github.com/JulyInSummer/cinematic/internal/app/storage"
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Logger *zap.Logger
	DB     *gorm.DB
	Config *Config
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

func NewStorage(logger *zap.Logger, conf *Config) (storage.MoviesI, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        getConnString(conf),
	}))
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		return nil, err
	}

	err = db.AutoMigrate(&models.Movie{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	return &Storage{
		Logger: logger,
		DB:     db,
		Config: conf,
	}, nil
}

func (s *Storage) Create(ctx context.Context, movie models.Movie) error {
	method := "Storage.Create"

	result := s.DB.WithContext(ctx).Create(&movie)
	if result.Error != nil {
		s.Logger.Error(method, zap.Error(result.Error))
		return result.Error
	}

	return nil
}

func (s *Storage) GetAll(ctx context.Context) ([]models.Movie, error) {
	method := "Storage.GetAll"

	var movies []models.Movie
	result := s.DB.WithContext(ctx).Find(&movies)
	if result.Error != nil {
		s.Logger.Error(method, zap.Error(result.Error))
		return nil, result.Error
	}

	return movies, nil
}

func (s *Storage) GetByID(ctx context.Context, id int) (*models.Movie, error) {
	method := "Storage.GetByID"

	var movie models.Movie
	result := s.DB.WithContext(ctx).First(&movie, id)
	if result.Error != nil {
		s.Logger.Error(method, zap.Error(result.Error))
		return nil, result.Error
	}

	return &movie, nil
}

func (s *Storage) Update(ctx context.Context, movie models.Movie) (*models.Movie, error) {
	method := "Storage.Update"

	result := s.DB.WithContext(ctx).Model(&movie).Updates(models.Movie{
		Title:    movie.Title,
		Plot:     movie.Plot,
		Director: movie.Director,
		Year:     movie.Year,
	})
	if result.Error != nil {
		s.Logger.Error(method, zap.Error(result.Error))
		return nil, result.Error
	}

	return &movie, nil
}

func (s *Storage) Delete(ctx context.Context, id int) error {
	method := "Storage.Delete"

	result := s.DB.WithContext(ctx).Delete(&models.Movie{}, id)
	if result.Error != nil {
		s.Logger.Error(method, zap.Error(result.Error))
		return result.Error
	}

	return nil
}
