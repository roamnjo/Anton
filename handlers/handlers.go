package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/roamnjo/Anton/data"
)

type Request struct {
	URL   string `json:"url"`
	Alias string `json:"alias"`
}

type Handler struct {
	storage *data.Storage
}

var req Request

func NewHandler(s *data.Storage) *Handler {
	return &Handler{storage: s}
}

func (h *Handler) PostUrl(rw http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error decoding req:", err)
		http.Error(rw, "Unexpected error", http.StatusInternalServerError)
		return
	}

	res, err := h.storage.SaveURL(req.URL, req.Alias)
	if err != nil {
		log.Println("Error saving url:", err)
		http.Error(rw, "Error saving data", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(res)
}

func (h *Handler) GetUrl(rw http.ResponseWriter, r *http.Request) {
	alias := r.URL.Query().Get("alias")

	url, err := h.storage.SelectURL(alias)
	if err != nil {
		log.Println("Error getting URL:", err)
		http.Error(rw, "Error getting data", http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]string{
		"alias": alias,
		"url":   url,
	})
}
