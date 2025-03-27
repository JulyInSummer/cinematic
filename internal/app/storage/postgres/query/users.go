package query

const (
	CreateUser = `INSERT INTO users (email, password, created_at) VALUES ($1, $2, now());`

	GetUserByEmail = `SELECT id, email, password FROM users WHERE email = $1 AND deleted_at IS NULL;`
)
