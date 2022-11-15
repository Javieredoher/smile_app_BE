package db

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

)


var (
	DB *sql.DB
)

func init() {

		//Enviroment
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error .env variables")
			return
		}
	
		host := os.Getenv("HOST")
		dbusername := os.Getenv("DB_USER")
		dbpass := os.Getenv("DB_PASS")
		dbname := os.Getenv("DB_NAME")
	
		// Database
		dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbusername, dbpass, host, dbname)
		db, err := sql.Open("mysql", dataSource)
		if err != nil {
			log.Fatal("error opening sql server")
			return
		}
	
		log.Println("database conected")
		DB = db
}
