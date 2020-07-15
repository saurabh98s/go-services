package handlers

import (
	"micro-services/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	// this will always be converted because of the router
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("[DEBUG] Handle DELETE Product", id)
	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product Not Found", http.StatusInternalServerError)
		return
	}
}
