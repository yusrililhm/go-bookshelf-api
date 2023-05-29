package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB) {

	user := GetEnv("USERS")
	pass := GetEnv("PASS")
	dbName := GetEnv("DBNAME")

	conn := user + ":" + pass + "@tcp(localhost:3306)/" + dbName
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return db
}

func GetEnv(k string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err.Error())
	}

	return os.Getenv(k)
}