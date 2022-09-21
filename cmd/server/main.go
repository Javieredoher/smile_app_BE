package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	StorageDB *sql.DB
)

func main() {

	//Enviroment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error .env variables")
		return
	}

	host := os.Getenv("HOST")
	dbusername := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// Database
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbusername, dbpass, host, dbport, dbname)
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("error opening sql server")
		return
	}

	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}

	log.Println("database conected")

	//Start server
	r := gin.Default()
	if err = r.Run(); err != nil {
		panic(err)
	}

	//router := routes.NewRouter(r, StorageDB)
	
}