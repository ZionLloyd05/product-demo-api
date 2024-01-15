package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for a Product API
type Product struct {
	ID 			int		`json:"id"`
	Name 		string	`json:"name"`
	Description string	`json:"description"`
	Price		float32	`json:"price"`
	SKU			string	`json:"sku"`
	CreatedOn	string	`json:"-"`
	UpdatedOn	string	`json:"-"`
	DeletedOn	string	`json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy Milky Coffee",
		Price: 2.45,
		SKU: "abc232",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Expresso",
		Description: "Short and strong coffee without milk",
		Price: 1.99,
		SKU: "fjd987",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}