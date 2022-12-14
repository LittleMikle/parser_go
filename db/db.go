package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectToDb() {
	dsn := "host=localhost port=5432 user=postgres password=123 dbname=postgres sslmode=disable timezone=Europe/Moscow"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
	}

	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Image{})

	Database = DbInstance{
		Db: db,
	}
}
