package database

const CreateSchema = `
CREATE TABLE IF NOT EXISTS users
(
	id SERIAL PRIMARY KEY NOT NULL,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	date_of_birth INTEGER,
	income_per_year NUMERIC(10, 2)
)
`

const InsertUserSchema = `
INSERT INTO users(first_name, last_name, date_of_birth, income_per_year) VALUES($1, $2, $3, $4) RETURNING id
`
