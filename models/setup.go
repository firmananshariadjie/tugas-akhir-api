package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin"))
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_restapi_gin_test?parseTime=true"))
	if err != nil {
		panic(err)
	}
	
	// Lakukan migrasi otomatis untuk tabel Event
	database.AutoMigrate(&Event{})
	// Lakukan migrasi otomatis untuk tabel Participant
	database.AutoMigrate(&Participant{})
	

	DB = database
}
