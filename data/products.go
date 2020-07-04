package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product holds the product data
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"` //omit from output
	UpdatedOn   string  `json:"-"` //omit from output
	DeletedOn   string  `json:"-"` //omit from output
}

// GetProducts returns arrays of products
func GetProducts() Products {
	return productList
}

// Products holds an array of products
type Products []*Product

// ToJSON Encodes data to JSON
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)

}

var productList = []*Product{

	{
		ID:          1,
		Name:        "Latte",
		Description: "frothy milk coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().Local().String(),
		UpdatedOn:   time.Now().Local().String(),
	},
	{

		ID:          1,
		Name:        "Espresso",
		Description: "Short and strong coffee , no milk",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().Local().String(),
		UpdatedOn:   time.Now().Local().String(),
	},
}
