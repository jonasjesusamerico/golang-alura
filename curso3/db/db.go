package db

import "database/sql"

func ConectaComBancoDados() *sql.DB {
	conexao := "user=postgres dbname=curso_alura password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
