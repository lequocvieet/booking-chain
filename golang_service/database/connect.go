package db

import (
	"golang_service/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=booking-chain port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Room{})
	db.AutoMigrate(&models.ListRoom{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.StateRoom{})
	db.AutoMigrate(&models.RoomNFT{})

	return db
}
