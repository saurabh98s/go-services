// Package classification of Product API
// 
// Documentation for Product APi
// 
// Schemes: http
// BasePath: /
// Version: 1.0.0
// 
// Consumes:
// -application/json
// 
// Produces:
// -application/json
// 
// swagger:meta


package handlers

import (
	"context"
	"fmt"
	"log"
	"micro-services/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products is a simple handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{
		l: l,
	}
}

// GetProducts returns the list of Products
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle GET Products")
	productList := data.GetProducts()
	err := productList.ToJSON(w) //encoding data from New Encoder
	if err != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

// AddProduct adds a new Product to the list
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST Products")
	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&product)

}

// UpdateProducts updates the value of a given product
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	// returns a map with the required value/
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.l.Printf("Handle PUT Products %d", id)
	product := r.Context().Value(KeyProduct{}).(data.Product)
	// Update value with corresponding data
	err = data.UpdateProduct(id, &product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

}

// KeyProduct stores the value of the context
type KeyProduct struct{}

// MiddlewareValidateProduct handles deserializing data
func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest)
			return
		}
		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
