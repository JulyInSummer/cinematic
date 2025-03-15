package domain

import (
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"
	"gorm.io/gorm"
)

type Movie struct {
	ID       uint
	Title    *string
	Director *string
	Year     *string
	Plot     *string
}

func (m *Movie) ToModel() models.Movie {
	return models.Movie{
		Model:    gorm.Model{ID: m.ID},
		Title:    m.Title,
		Director: m.Director,
		Year:     m.Year,
		Plot:     m.Plot,
	}
}
