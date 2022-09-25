package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)


var (
	DB *gorm.DB
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
		dbport := os.Getenv("DB_PORT")
		dbname := os.Getenv("DB_NAME")
	
		// Database
		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbusername, dbpass, host, dbport, dbname)
		//db, err := sql.Open("mysql", dataSource)
		db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
		if err != nil {
			log.Fatal("error opening sql server")
			return
		}
	
		log.Println("database conected")
		DB = db
}
