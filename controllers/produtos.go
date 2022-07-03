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

		precoConvertidoParaFloat := converterParaFloat64(preco)
		quantidadeConvertidoParaInt := converterParaInt(quantidade)

		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)

	}
}

func converterParaFloat64(valor string) float64 {
	valorConvertido, err := strconv.ParseFloat(valor, 64)

	if err != nil {
		log.Println("Erro na conversão para float64", err)
	}

	return valorConvertido
}

func converterParaInt(valor string) int {
	valorConvertido, err := strconv.Atoi(valor)

	if err != nil {
		log.Println("Erro na conversão para inteiro", err)
	}

	return valorConvertido
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	produto := models.BuscarProduto(idProduto)

	templates.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat := converterParaFloat64(preco)
		quantidadeConvertidoParaInt := converterParaInt(quantidade)
		idConvertidoParaInt := converterParaInt(id)

		models.EditarProduto(idConvertidoParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)

	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
