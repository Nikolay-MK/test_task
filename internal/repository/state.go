package repository

import (
	"fmt"
)

// isDataReady проверка наличия данных
func (r *Repository) IsDataReady() bool {

	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM individuals").Scan(&count)
	if err != nil {
		fmt.Println("Error checking data in the database:", err)
		return false
	}

	return count > 0
}
