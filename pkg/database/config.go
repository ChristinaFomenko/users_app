package database

import "fmt"

var (
	dbUsername = "usersappdb"
	dbPassword = "userasappdb"
	dbHost     = "localhost"
	dbTable    = "usersappdb"
	dbPort     = "5432"
	pgConnStr  = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
)
