package repository

const (
	queryInsertUser = `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			phone,
			created_at,
			updated_at,
			user_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	queryGetAll = `SELECT * FROM users`
	queryGetOne = `SELECT * FROM users WHERE user_id=$1;`
)
