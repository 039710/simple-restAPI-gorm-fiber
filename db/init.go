package db

import (
	"fmt"
	"os"

	model "go-learn/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// use ENV variables to connect to DB
	fmt.Println("Connecting to DB...")
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	dsn := os.Getenv("DB_DSN")
	// parse the DSN string
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
	fmt.Println("Connected to DB")
	// drop table if exist
	db.Migrator().DropTable(&model.Student{}, &model.Subject{}, &model.StudentSubject{})
	db.AutoMigrate(&model.Student{}, &model.Subject{}, &model.StudentSubject{})

	DB = db
}
