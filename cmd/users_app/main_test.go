package main

//import (
//	"log"
//	"os"
//	"testing"
//
//	"github.com/ChristinaFomenko/users_app/pkg/database"
//
//	"github.com/jmoiron/sqlx"
//)
//
//const (
//	dbDriver = "postgres"
//	dbSource = "postgresql://usersappdb:usersappdb@localhost:5432/usersappd?sslmode=disable"
//)
//
//var testUserRepository *database.UserRepository
//
////var testNewApp *handler.App
//
//func TestMain(m *testing.M) {
//	conn, err := sqlx.Open(dbDriver, dbSource)
//	if err != nil {
//		log.Fatal("cannot connect to db:", err)
//	}
//	testUserRepository = database.NewUserRepository(conn)
//	//testNewApp = handler.NewApp(conn)
//
//	os.Exit(m.Run())
//}
