package handlers

import (
	"log"
	"net/http"

	"github.com/binjamil/working/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
	} else if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// getProducts returns the products from our data store
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// Fetch product list from the data store
	lp := data.GetProducts()

	// Serialize the list to JSON and write response
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// addProduct creates a new product
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	// Create a new product from JSON
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse json", http.StatusBadRequest)
	}

	// Persist the new product to our data store
	data.AddProduct(prod)
}
