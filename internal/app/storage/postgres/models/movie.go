package models

type Movie struct {
	ID       uint64
	Title    *string
	Director *string
	Year     *string
	Plot     *string
}
