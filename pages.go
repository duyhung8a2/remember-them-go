package main

import "net/http"

type PageHandler struct {
}

func (b PageHandler) ListPages(w http.ResponseWriter, r *http.Request)  {}
func (b PageHandler) GetPages(w http.ResponseWriter, r *http.Request)   {}
func (b PageHandler) CreatePage(w http.ResponseWriter, r *http.Request) {}
func (b PageHandler) UpdatePage(w http.ResponseWriter, r *http.Request) {}
func (b PageHandler) DeletePage(w http.ResponseWriter, r *http.Request) {}