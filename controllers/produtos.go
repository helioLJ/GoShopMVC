package controllers

import (
	"net/http"
	"productsAlura/models"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		precoStr := r.FormValue("preco")
		preco, err := strconv.ParseFloat(precoStr, 64)
		if err != nil {
			http.Error(w, "Erro ao converter preço", http.StatusBadRequest)
			return
		}
		quantidadeStr := r.FormValue("quantidade")
		quantidade, err := strconv.Atoi(quantidadeStr)
		if err != nil {
			http.Error(w, "Erro ao converter quantidade", http.StatusBadRequest)
			return
		}
		models.CriaNovoProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProduto)
	temp.ExecuteTemplate(w, "edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idProduto := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		precoStr := r.FormValue("preco")
		preco, err := strconv.ParseFloat(precoStr, 64)
		if err != nil {
			http.Error(w, "Erro ao converter preço", http.StatusBadRequest)
			return
		}
		quantidadeStr := r.FormValue("quantidade")
		quantidade, err := strconv.Atoi(quantidadeStr)
		if err != nil {
			http.Error(w, "Erro ao converter quantidade", http.StatusBadRequest)
			return
		}
		models.AtualizaProduto(idProduto, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}