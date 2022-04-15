package main

import (
	"golang-alura/curso3/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
