package main

import (
	"context"
	"encoding/json"
	"net/http"
	"remember_them/models"

	"github.com/stephenafamo/bob"
)

type PageHandler struct {
	db *bob.DB
}

func NewPageHandler(db *bob.DB) *PageHandler {
	return &PageHandler{db: db}
}

func (b PageHandler) ListPages(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	pageTable := models.Pages

	pages, err := pageTable.Query(ctx, b.db).All()
	if err != nil {
		http.Error(w, "Failed to retrieve pages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pages)
}
func (b PageHandler) GetPages(w http.ResponseWriter, r *http.Request) {

}
func (b PageHandler) CreatePage(w http.ResponseWriter, r *http.Request) {}
func (b PageHandler) UpdatePage(w http.ResponseWriter, r *http.Request) {}
func (b PageHandler) DeletePage(w http.ResponseWriter, r *http.Request) {}
