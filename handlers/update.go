package handlers

import (
	"micro-services/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
