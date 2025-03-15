package request

import (
	"github.com/JulyInSummer/cinematic/internal/app/service/domain"
)

type CreateMovie struct {
	Title    string `json:"title" binding:"required"`
	Year     string `json:"year" binding:"required,date_format"`
	Director string `json:"director" binding:"required"`
	Plot     string `json:"plot" binding:"required"`
}

func (r CreateMovie) ToDomain() domain.Movie {
	return domain.Movie{
		Title:    &r.Title,
		Director: &r.Director,
		Year:     &r.Year,
		Plot:     &r.Plot,
	}
}

type UpdateMovie struct {
	ID       uint    `json:"id"`
	Title    *string `json:"title"`
	Year     *string `json:"year"`
	Plot     *string `json:"plot"`
	Director *string `json:"director"`
}

func (r UpdateMovie) ToDomain() domain.Movie {
	return domain.Movie{
		ID:       r.ID,
		Title:    r.Title,
		Year:     r.Year,
		Plot:     r.Plot,
		Director: r.Director,
	}
}
