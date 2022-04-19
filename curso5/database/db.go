package database

import (
	"golang-alura/curso5/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDados() {
	dsn := "host=localhost user=postgres password=postgres dbname=alura port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
