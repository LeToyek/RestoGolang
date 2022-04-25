package repository

const (
	queryInsertUser = `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			phone,
			token,
			refresh_token,
			user_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
)
