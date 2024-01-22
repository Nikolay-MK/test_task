package repository

import "selltech/internal/selltech"

// UpdateIndividuals обновляет таблицу individuals в базе данных PostgreSQL
func (r *Repository) UpdateIndividuals(entries []selltech.SDNEntry) error {
	for _, entry := range entries {
		_, err := r.db.Exec(`INSERT INTO individuals (first_name, last_name)
		VALUES ($1, $2)
		ON CONFLICT (first_name, last_name) DO NOTHING `,
			//SET first_name = $1, last_name = $2;`,
			entry.FirstName, entry.LastName)
		if err != nil {
			return err
		}
	}
	return nil
}
