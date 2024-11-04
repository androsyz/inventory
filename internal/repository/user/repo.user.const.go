package user

const (
	sqlInsertUser = `
		INSERT INTO users (name) 
		VALUES ($1)
		RETURNING id;
	`

	sqlGetUsers = `
		SELECT id, name
		FROM users
		ORDER BY id DESC;
	`
)
