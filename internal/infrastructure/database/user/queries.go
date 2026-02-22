package user

var (
	GetUserByIDQuery = `SELECT id,first_name,last_name,dni,email,telephone,created_at,updated_at FROM users WHERE id=$1;`
)
