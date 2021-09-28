package database

const createSchema = `
CREATE TABLE IF NOT EXISTS users
( 
	id SERIAL PRIMARY KEY,
	first_name TEXT,
	last_name TEXT,
	date_of_birth TEXT,
	income_per_year TEXT
)
`

var insertUserSchema = `
INSERT INTO users(first_name, last_name, date_of_birth, income_per_year) VALUES($1, $2, $3, $4) RETURNING id
`
