package repository

import (
	"github.com/lib/pq"
	"selltech/internal/selltech"
)

// UpdateIndividuals обновляет таблицу individuals в базе данных
func (r *Repository) UpdateIndividuals(entries []selltech.SDNEntry) error {
	var individuals []selltech.SDNEntry
	for _, entry := range entries {
		if entry.SDNType == "Individual" {
			individuals = append(individuals, entry)
		}
	}

	if len(individuals) == 0 {
		// Нет записей для вставки
		return nil
	}

	// Подготавливаем запрос
	stmt := `INSERT INTO individuals (uid, first_name, last_name)
             SELECT unnest($1::int[]), unnest($2::text[]), unnest($3::text[])
             ON CONFLICT (uid) DO UPDATE 
             SET first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name;`

	//подготавливаем аргументы для передачи в Exec
	var uidArray, firstNameArray, lastNameArray pq.StringArray
	for _, ind := range individuals {
		uidArray = append(uidArray, ind.Uid)
		firstNameArray = append(firstNameArray, ind.FirstName)
		lastNameArray = append(lastNameArray, ind.LastName)
	}

	// выполняем запрос в бд
	_, err := r.db.Exec(stmt, uidArray, firstNameArray, lastNameArray)
	if err != nil {
		return err
	}

	return nil
}
