package main

import (
	"go-rest-api/common"
	"go-rest-api/internal/models"
	"log"
)

func main() {
	db, err := common.NewMySql()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	log.Println("Migration up successfully")
}
