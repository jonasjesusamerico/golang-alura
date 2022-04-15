package models

import "golang-alura/curso3/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDados()

	selectTodosProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosProdutos.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int

		err = selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade ) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(idProduto string) {
	db := db.ConectaComBancoDados()

	deletaOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletaOProduto.Exec(idProduto)

	defer deletaOProduto.Close()

}

func EditaProduto(idProduto string) Produto {
	db := db.ConectaComBancoDados()

	produtoDobanco, err := db.Query("select * from produtos where id=$1", idProduto)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualiza := Produto{}

	for produtoDobanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := produtoDobanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualiza.Id = id
		produtoParaAtualiza.Nome = nome
		produtoParaAtualiza.Descricao = descricao
		produtoParaAtualiza.Preco = preco
		produtoParaAtualiza.Quantidade = quantidade

	}

	defer produtoDobanco.Close()
	return produtoParaAtualiza
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDados()

	atualizaproduto, err := db.Prepare("update produtos set nome=$2, descricao=$3, preco=$4, quantidade=$5 where id=$1")
	if err != nil {
		panic(err.Error())
	}

	atualizaproduto.Exec(id, nome, descricao, preco, quantidade)
	defer db.Close()
}
