package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title    *string
	Director *string
	Year     *string
	Plot     *string
}
