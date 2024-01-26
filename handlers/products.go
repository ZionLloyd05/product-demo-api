package handlers

import (
	"log"
	"main/data"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	
	// Get all
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// Post
	if r.Method == http.MethodPost {
		p.addProoduct(rw, r)
		return
	}

	// Update
    if r.Method == http.MethodPut {
		rg := regexp.MustCompile(`/([0-9]+)`)
		group := rg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, rw, r)
		return 
	}
	

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle Get Products")

	productList := data.GetProducts()

	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
	}
}

func (p *Products) addProoduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle Post Product")

	prod := &data.Product{}
	
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle Update Product")

	prod := &data.Product{}
	
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}