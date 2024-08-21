package database

import (
	"backend_technical_test/entities/models"
	"fmt"
	"log"
)

func DBMigrate() {
	err := DB.AutoMigrate(&models.Posts{})

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Migration Success")
}
