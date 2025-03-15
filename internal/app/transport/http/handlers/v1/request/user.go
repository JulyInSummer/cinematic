package request

import "github.com/JulyInSummer/cinematic/internal/app/service/domain"

type CreateUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (r CreateUser) ToDomain() domain.User {
	return domain.User{
		Email:    r.Email,
		Password: r.Password,
	}
}
