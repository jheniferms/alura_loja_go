package main

import (
	"net/http"

	"github.com/jheniferms/alura_loja_go/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
