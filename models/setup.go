package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Gagal melakukan load env")
	}
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:"+DBPort+")/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DBHost, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database 111" + DbUser)
	}
	db.AutoMigrate(&User{}, &Product{})
	DB = db
}
