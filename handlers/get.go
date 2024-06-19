package handlers

import (
	"net/http"
	"remember_them/data"
)

// swagger:route GET /products products listProducts
// Return a list of products
// Responses:
//
//	200: productsResponse
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
