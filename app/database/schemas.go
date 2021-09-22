package database

const createSchema = `
CREATE TABLE IF NOT EXISTS users
( 
	id SERIAL PRIMARY KEY,
	first_name TEXT,
	last_name TEXT,
	date_of_birth DATE
	income_per_year NUMERIC(10, 2)
)
`
