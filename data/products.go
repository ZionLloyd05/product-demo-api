package data

import (
	"encoding/json"
	"fmt"
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

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()

	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList) - 1]
	var currentId = lp.ID

	return currentId + 1
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