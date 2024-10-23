package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	UpdatedOn   string  `json:"updatedOn"`
	CreatedOn   string  `json:"createOn"`
	DeletedOn   string  `json:"deletedOn"`
}

func GetProducts() Products {
	return productList
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.Id = getNextId()
	productList = append(productList, p)
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.Id + 1
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.Id = id
	productList[pos] = p

	return nil

}

var ErrorProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.Id == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

var productList = []*Product{
	{
		Id:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		Id:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
