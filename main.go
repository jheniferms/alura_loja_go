package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	produtosFromDatabase, err := db.Query("select * from produtos")

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

	templates.ExecuteTemplate(w, "Index", produtos)

	defer db.Close()
}
