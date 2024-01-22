package repository

import (
	"fmt"
	"selltech/internal/selltech"
)

// GetNames возвращает имена в зависимости от условий запроса
func (r *Repository) GetNames(query string, args ...interface{}) []selltech.NameResult {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return nil
	}
	defer rows.Close()

	var result []selltech.NameResult
	for rows.Next() {
		var nameResult selltech.NameResult
		if err := rows.Scan(&nameResult.UID, &nameResult.FirstName, &nameResult.LastName); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil
		}
		result = append(result, nameResult)
	}

	return result
}

// GetStrongNames возвращает имена с точным совпадением
func (r *Repository) GetStrongNames(name string) []selltech.NameResult {
	query := "SELECT uid, first_name, last_name FROM individuals WHERE first_name = $1 OR last_name = $1"
	return r.GetNames(query, name)
}

// GetWeakNames возвращает имена с любым совпадением
func (r *Repository) GetWeakNames(name string) []selltech.NameResult {
	query := "SELECT uid, first_name, last_name FROM individuals WHERE first_name ILIKE $1 OR last_name ILIKE $1"
	return r.GetNames(query, "%"+name+"%")
}

// GetAllNames возвращает все имена
func (r *Repository) GetAllNames(name string) []selltech.NameResult {
	query := "SELECT uid, first_name, last_name FROM individuals"
	return r.GetNames(query)
}
