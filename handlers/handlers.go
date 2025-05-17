package handlers

import (
	"database/sql"

	"github.com/boka18/repartners-interview/calculator"
)

func getPackSizesFromDB(db *sql.DB) ([]calculator.Pack, error) {
	rows, err := db.Query("SELECT id, size FROM pack_sizes ORDER BY size ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sizes []calculator.Pack
	for rows.Next() {
		var id int
		var size int
		if err := rows.Scan(&id, &size); err != nil {
			return nil, err
		}
		sizes = append(sizes, calculator.Pack{
			ID:   id,
			Size: size,
		})
	}
	return sizes, nil
}

func deletePackSizesFromDB(db *sql.DB, id uint) error {
	_, err := db.Exec("DELETE FROM pack_sizes WHERE id = $1", id)
	return err
}

func addPackSizeToDB(db *sql.DB, size int) error {
	_, err := db.Exec(
		"INSERT INTO pack_sizes (size) VALUES ($1)",
		size,
	)
	return err
}
