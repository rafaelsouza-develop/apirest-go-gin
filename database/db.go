package database

import (
	"apirest-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Deu ruim ao conectar")
	}
	DB.AutoMigrate(&models.Aluno{})

}
