package main

import (
	"golang-alura/curso5/database"
	"golang-alura/curso5/routes"
)

func main() {
	database.ConectaComBancoDados()
	routes.HandleRequest()
}
