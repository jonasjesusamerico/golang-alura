package main

import (
	"fmt"
	"golang-alura/curso4/database"
	"golang-alura/curso4/models"
	"golang-alura/curso4/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "historia 2"},
	}

	database.ConectaComBancoDados()
	fmt.Println("Iniciando o servidor rest com o GO!")
	routes.HandeRequest()
}
