package database

import (
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := "postgresql://authdb_owner:y1af2TerlsRU@ep-weathered-voice-a5n7l308.us-east-2.aws.neon.tech/authdb?sslmode=require"

	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db = connect
	
}



func GetDB() *gorm.DB {
	return db
}