package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/boka18/repartners-interview/calculator"
)

type ApiResponseCalculate struct {
	TotalItems int         `json:"total_items"`
	PacksUsed  map[int]int `json:"packs_used"`
}

// RegisterCalculateHandler attaches the GET /calculate handler to the router
func RegisterCalculateHandler(mux *http.ServeMux, db *sql.DB) {
	mux.HandleFunc("/api/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		calculator := calculator.NewPackSize()
		orderStr := r.URL.Query().Get("order")
		if orderStr == "" {
			http.Error(w, "Missing 'order' query parameter", http.StatusBadRequest)
			return
		}

		order, err := strconv.Atoi(orderStr)
		if err != nil {
			http.Error(w, "Invalid 'order' value", http.StatusBadRequest)
			return
		}

		// Read pack sizes from DB
		packSizes, err := getPackSizesFromDB(db)
		if err != nil {
			http.Error(w, "Failed to fetch pack sizes", http.StatusInternalServerError)
			return
		}

		packSizesByValue := make([]int, len(packSizes))
		for i, packSize := range packSizes {
			packSizesByValue[i] = packSize.Size
		}
		result := calculator.Calculate(packSizesByValue, order)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ApiResponseCalculate{
			TotalItems: result.TotalItems,
			PacksUsed:  result.PacksUsed,
		})
	})
}
