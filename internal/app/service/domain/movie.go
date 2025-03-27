package domain

import (
	"github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"
)

type Movie struct {
	ID       uint64
	Title    *string
	Director *string
	Year     *string
	Plot     *string
}

func (m *Movie) ToModel() models.Movie {
	return models.Movie{
		ID:       m.ID,
		Title:    m.Title,
		Director: m.Director,
		Year:     m.Year,
		Plot:     m.Plot,
	}
}
