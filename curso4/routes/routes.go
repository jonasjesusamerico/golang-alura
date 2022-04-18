package routes

import (
	"golang-alura/curso4/controllers"
	"golang-alura/curso4/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandeRequest() {

	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleWare)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
