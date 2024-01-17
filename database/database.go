package database

import (
	"stury-supplements/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComDB() {
	stringConnection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection))
	var suplementos models.Suplementos
	var user models.User
	DB.AutoMigrate(&suplementos, &user)

}
