package handlers

import (
	"log"
	"micro-services/data"
	"net/http"
)

// Products is a simple handler
type Products struct {
	l *log.Logger
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
	}
	// Catch rest of the request
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle GET Products")
	productList := data.GetProducts()
	err := productList.ToJSON(w) //encoding data from New Encoder
	if err != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle POST Products")
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to UnMarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(product)

}

// NewProducts creates a new products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{
		l: l,
	}
}
