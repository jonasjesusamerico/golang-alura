package main

import (
	"golang-alura/curso6/database"
	"golang-alura/curso6/routes"
)

func main() {
	database.ConectaComBancoDados()
	routes.HandleRequest()
}
