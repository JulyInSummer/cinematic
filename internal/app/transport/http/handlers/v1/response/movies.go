package response

type Movie struct {
	ID       uint64  `json:"id"`
	Title    *string `json:"title"`
	Director *string `json:"director"`
	Year     *string `json:"year"`
	Plot     *string `json:"plot"`
}
