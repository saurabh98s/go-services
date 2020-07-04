package handlers

import (
	"encoding/json"
	"log"
	"micro-services/data"
	"net/http"
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

func (p *Products)ServeHTTP(w http.ResponseWriter, h *http.Request) {
	 productList:=data.GetProducts()
	 data,err:=json.Marshal(productList)
	 if err != nil {
		 http.Error(w,"Unable to Marshal JSON",http.StatusInternalServerError)
	 }
	 w.Write(data)
}