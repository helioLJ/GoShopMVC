package main

import (
	"net/http"
	"productsAlura/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}