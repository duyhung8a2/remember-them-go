//	Package classification of Product API
//
//	Documentation for Product API
//
//	Schemes: http, https
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"remember_them/data"

	"github.com/go-chi/chi/v5"
)

// A list of products return in the response
// swagger:response productsResponse
type ProductsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product 
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", p.getProducts)
	r.With(p.MiddlewareProductValidation).Post("/", p.addProduct)
	r.With(p.MiddlewareProductValidation).Put("/{id}", p.updateProducts)
	return r
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		}

		// Validate product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// Add product to context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
