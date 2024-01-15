package handlers

import (
	"log"
	"main/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request){
	
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request){
	productList := data.GetProducts()

	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
	}
}