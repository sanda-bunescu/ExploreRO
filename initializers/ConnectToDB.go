package initializers

import (
	"log"
	"os"

	"github.com/sanda-bunescu/ExploreRO/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = db
}

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}

