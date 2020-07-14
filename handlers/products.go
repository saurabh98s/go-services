// Package handlers of Product API
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
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in:body
	Body []data.Product
}

// Products is a http handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{
		l: l,
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
