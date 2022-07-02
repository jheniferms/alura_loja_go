package models

import "github.com/jheniferms/alura_loja_go/db2"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db2.ConectaComBancoDeDados()

	produtosFromDatabase, err := db.Query("select * from produtos order by nome")

	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}
	for produtosFromDatabase.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtosFromDatabase.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produto := Produto{id, nome, descricao, preco, quantidade}

		produtos = append(produtos, produto)

	}

	defer db.Close()
	return produtos

}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db2.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db2.ConectaComBancoDeDados()

	deletarProdutoQuery, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarProdutoQuery.Exec(id)

	defer db.Close()
}

func BuscarProduto(id string) Produto {
	db := db2.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoRetorno := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoRetorno.Nome = nome
		produtoRetorno.Descricao = descricao
		produtoRetorno.Quantidade = quantidade
		produtoRetorno.Preco = preco
		produtoRetorno.Id = id
	}

	defer db.Close()

	return produtoRetorno
}

func EditarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db2.ConectaComBancoDeDados()

	updateQuery, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateQuery.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
