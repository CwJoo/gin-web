package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//Setup init database connection
func Setup() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/ginweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	Migrate()
	return
}

//Migrate used to migate database changes
func Migrate() {
	db.AutoMigrate(&User{})
}
