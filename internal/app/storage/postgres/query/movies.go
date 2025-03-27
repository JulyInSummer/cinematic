package query

const (
	CreateMovie = `INSERT INTO movies (title, director, year, plot, created_at) values ($1, $2, $3, $4, now())`

	GetAllMovies = `SELECT id, title, director, year, plot FROM movies WHERE deleted_at IS NULL ORDER BY title`

	GetMovieByID = `SELECT id, title, director, year, plot FROM movies WHERE id = $1 and deleted_at IS NULL`

	UpdateMovie = `UPDATE movies SET title = coalesce(@title, title), director = coalesce(@director, director), year = coalesce(@year, year), plot = coalesce(@plot, plot), updated_at = now() WHERE id = @id;`

	DeleteMovie = `UPDATE movies SET deleted_at = now() WHERE id = $1;`
)
