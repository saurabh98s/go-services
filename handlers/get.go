package handlers

import (
	"micro-services/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of Products
// Responses:
// 200: productResponse

// GetProducts returns the list of Products
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle GET Products")
	productList := data.GetProducts()
	err := productList.ToJSON(w) //encoding data from New Encoder
	if err != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}
