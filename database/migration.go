package database

import (
	"be-waybucks/models"
	"be-waybucks/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Transaction{},
		&models.Profile{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
