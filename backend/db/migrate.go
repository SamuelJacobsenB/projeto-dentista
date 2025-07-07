package db

import (
	"fmt"

	"github.com/SamuelJacobsenB/projeto-dentista/entities"
)

func Migrate() {
	err := DB.AutoMigrate(&entities.User{}, &entities.Patient{})

	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	fmt.Println("Database migrated successfully")
}
