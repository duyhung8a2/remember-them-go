package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Get data from body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Fatal(err)
	}

	// First create a struct or map to hold JSON data
	// Struct
	type Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		Data    string `json:"data"`
	}
	payload := Response{
		Message: "OK",
		Status:  "success",
		Data:    string(data),
	}
	h.l.Printf("Payload: %+v\n", payload)

	// Map
	// data := map[string]string{
	// 	"message": "OK",
	// 	"status":  "success",
	// }

	// Set content header to application/json
	rw.Header().Set("Content-Type", "application/json")

	// Marshal the data into JSON format
	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// okByte := []byte("OK")
	rw.Write(jsonData)
}
