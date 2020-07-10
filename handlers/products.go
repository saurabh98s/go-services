package handlers

import (
	"context"
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

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodGet {
	// 	p.getProducts(w, r)
	// 	return
	// }

	// if r.Method == http.MethodPost {
	// 	p.addProduct(w, r)
	// }

	// if r.Method == http.MethodPut {
	// 	p.l.Println("PUT")
	// 	// expecting id in the URI
	// 	path := r.URL.Path
	// 	reg := regexp.MustCompile(`/([0-9]+)`)
	// 	group := reg.FindAllStringSubmatch(path, -1)

	// 	if len(group) != 1 {
	// 		http.Error(w, "Invalid URL", http.StatusBadRequest)
	// 		return
	// 	}

	// 	if len(group) >= 2 {
	// 		p.l.Println("More than one capture group")
	// 		http.Error(w, "Invalid URL", http.StatusBadRequest)
	// 		return
	// 	}

	// 	idString := group[0][1]
	// 	id, err := strconv.Atoi(idString)
	// 	if err != nil {
	// 		http.Error(w, "Invalid URL", http.StatusBadRequest)
	// 		return
	// 	}
	// 	p.updateProducts(id, w, r)
	// 	return

	// }
	// // Catch rest of the request
	// w.WriteHeader(http.StatusMethodNotAllowed)
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
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to UnMarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(product)

}

// UpdateProducts updates the value of a given product
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.l.Printf("Handle PUT Products %d", id)
	//  product := &data.Product{}
	product := r.Context().Value(KeyProduct{}).(data.Product)
	// err = product.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(w, "Unable to UnMarshal JSON", http.StatusBadRequest)
	// }
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

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
