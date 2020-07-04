package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}


// ServeHTTP implements the go http.Handler interface (https://golang.org/pkg/net/http/#Handler)
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello World")

	// read the body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "OOps", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(w, "Hello %s", d)

}
