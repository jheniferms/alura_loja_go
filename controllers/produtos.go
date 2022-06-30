package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/jheniferms/alura_loja_go/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		quantidade := r.FormValue("quantidade")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão do quantidade", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)

		http.Redirect(w, r, "/", 301)

	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}
