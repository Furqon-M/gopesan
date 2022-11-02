package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/pesan"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Pesandb{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Menu{})
	DB = db
}
