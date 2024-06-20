package database


import (
	"log"

	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"

)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/shalomdb?charset=utf8&parseTime=True&loc=Local"

	connect, err := gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db = connect
}



func GetDB() *gorm.DB {
	return db
}