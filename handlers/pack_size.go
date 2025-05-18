package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// RegisterCalculateHandler attaches the GET /calculate handler to the router
func RegisterPackSizeHandler(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/api/pack-size/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

		idStr := strings.TrimPrefix(r.URL.Path, "/api/pack-size/")
		if idStr == "" {
			http.Error(w, "Missing ID", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		handlePackSizeDel(w, r, db, id)
	})

	mux.HandleFunc("/api/pack-size", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlePackSizeGet(w, r, db)
		case http.MethodPost:
			handlePackSizePost(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func handlePackSizeGet(w http.ResponseWriter, _ *http.Request, db *sql.DB) {
	res, err := getPackSizesFromDB(db)
	if err != nil {
		http.Error(w, "Failed to fetch pack sizes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func handlePackSizePost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	type requestData struct {
		Value int `json:"value"`
	}

	var data requestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	valueInt := data.Value
	err = addPackSizeToDB(db, valueInt)
	if err != nil {
		http.Error(w, "Failed to add pack size", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{message: ok}`)
}

func handlePackSizeDel(w http.ResponseWriter, _ *http.Request, db *sql.DB, id int) {
	err := deletePackSizesFromDB(db, uint(id))
	if err != nil {
		http.Error(w, "Failed to delete pack size", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{message: ok}`)
}
