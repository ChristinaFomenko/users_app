package database

const createSchema = `
CREATE TABLE IF NOT EXISTS users
( 
	id SERIAL PRIMARY KEY NOT NULL,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	date_of_birth TEXT,
	income_per_year NUMERIC(10, 2)
)
`

var insertUserSchema = `
INSERT INTO users(first_name, last_name, date_of_birth, income_per_year) VALUES($1, $2, $3, $4) RETURNING id
`
