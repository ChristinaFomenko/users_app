package database

import "fmt"

//TODO: это должны быть константы, а не просто переменные.
var (
	dbUsername = "usersappdb"
	dbPassword = "userasappdb"
	dbHost     = "localhost"
	dbTable    = "usersappdb"
	dbPort     = "5432"
	//TODO: pgConnStr - строка, которая состоит из констант, она здесь не должна быть
	pgConnStr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
)
