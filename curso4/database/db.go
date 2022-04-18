package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDados() {
	// import "gorm.io/driver/postgres"
	// ref: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	dsn := "host=localhost user=postgres password=postgres dbname=alura port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}

}
