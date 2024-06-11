package api

import (
	"encoding/json"
	"net/http"
	"remember_them/models"
	"strconv"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
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
	ctx := r.Context()
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
	err = json.NewEncoder(w).Encode(pageResponses)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
func (b PageHandler) GetPages(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		http.Error(w, "Id is not an integer", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	// page, err := models.Pages.Query(ctx, b.db, sm.Where(models.PageColumns.ID.EQ(sqlite.Arg(id)))).One()
	page, err := models.FindPage(ctx, b.db, int32(id))
	if err != nil {
		http.Error(w, "Failed to retrieve pages", http.StatusInternalServerError)
		return
	}
	if page == nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	pageResponse := PageResponse{
		ID:        page.ID,
		Title:     page.Title,
		UserID:    page.UserID,
		ParentID:  page.ParentID,
		CreatedAt: page.CreatedAt,
		UpdatedAt: page.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(pageResponse)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

type CreatePageRequest struct {
	Title    string              `json:"title" validate:"required"`
	UserID   int32               `json:"user_id" validate:"required"`
	ParentID omitnull.Val[int32] `json:"parent_id"`
}

func (b PageHandler) CreatePage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreatePageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Not checking relation, i don't know why
	page, err := models.Pages.Insert(ctx, b.db, &models.PageSetter{
		Title:    omit.From(req.Title),
		UserID:   omit.From(req.UserID),
		ParentID: req.ParentID,
	})
	if err != nil {
		http.Error(w, "Failed to create page", http.StatusInternalServerError)
		return
	}

	pageResponse := PageResponse{
		ID:        page.ID,
		Title:     page.Title,
		UserID:    page.UserID,
		ParentID:  page.ParentID,
		CreatedAt: page.CreatedAt,
		UpdatedAt: page.CreatedAt,
	}
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(pageResponse)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

type UpdatePageRequest struct {
	Title    string              `json:"title" validate:"required"`
	UserID   int32               `json:"user_id" validate:"required"`
	ParentID omitnull.Val[int32] `json:"parent_id"`
}

func (b PageHandler) UpdatePage(w http.ResponseWriter, r *http.Request) {

}
func (b PageHandler) DeletePage(w http.ResponseWriter, r *http.Request) {}
