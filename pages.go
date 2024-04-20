package main

import (
	"context"
	"encoding/json"
	"net/http"
	"remember_them/models"
	"time"

	"github.com/stephenafamo/bob"
)

type PageHandler struct {
	db *bob.DB
}

func NewPageHandler(db *bob.DB) *PageHandler {
	return &PageHandler{db: db}
}

type PageResponse struct {
	ID        int32       `json:"id"`
	Title     string      `json:"title"`
	UserID    int32       `json:"user_id"`
	ParentID  interface{} `json:"parent_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (b PageHandler) ListPages(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	pageTable := models.Pages

	pages, err := pageTable.Query(ctx, b.db).All()
	if err != nil {
		http.Error(w, "Failed to retrieve pages", http.StatusInternalServerError)
		return
	}

	var pageResponses []PageResponse
	for _, page := range pages {
		pageResponse := PageResponse{
			ID:        page.ID,
			Title:     page.Title,
			UserID:    page.UserID,
			ParentID:  page.ParentID,
			CreatedAt: page.CreatedAt,
			UpdatedAt: page.CreatedAt,
		}
		pageResponses = append(pageResponses, pageResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pageResponses)
}
func (b PageHandler) GetPages(w http.ResponseWriter, r *http.Request)   {}
func (b PageHandler) CreatePage(w http.ResponseWriter, r *http.Request) {}
func (b PageHandler) UpdatePage(w http.ResponseWriter, r *http.Request) {}
func (b PageHandler) DeletePage(w http.ResponseWriter, r *http.Request) {}
