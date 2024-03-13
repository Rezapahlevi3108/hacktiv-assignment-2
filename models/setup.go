package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/hacktiv_assignment_2"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Item{}, &Order{})

	DB = database
}
