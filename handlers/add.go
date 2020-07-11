package handlers

import (
	"micro-services/data"
	"net/http"
)

// AddProduct adds a new Product to the list
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST Products")
	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&product)

}
