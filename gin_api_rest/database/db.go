package database

import (
	"api_go_gin/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {

	const (
		host     = "localhost"
		port     = 5432
		user     = "root"
		password = "root"
		dbname   = "root"
	)

	connectionUrl := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Student{})

	return db
}
