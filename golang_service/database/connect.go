package db

import (
	"fmt"
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

	err = db.AutoMigrate(&models.Room{})
	fmt.Print(err)
	db.AutoMigrate(&models.ListRoom{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.StateRoom{})
	db.AutoMigrate(&models.RoomNFT{})

	return db
}
