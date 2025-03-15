package domain

import "github.com/JulyInSummer/cinematic/internal/app/storage/postgres/models"

type User struct {
	Email    string
	Password string
}

func (d User) ToModel() models.User {
	return models.User{
		Email:    d.Email,
		Password: d.Password,
	}
}

type Token struct {
	Token string
}
