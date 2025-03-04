package database

import (
	"github.com/lans97/cassist-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects to the database
func InitDB() {
	dsn := "host=db user=myuser password=mypassword dbname=cassist port=5432 sslmode=disable TimeZone=America/Mexico_City"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	DB.AutoMigrate(&models.User{})
}
