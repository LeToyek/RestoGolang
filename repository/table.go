package repository

import (
	"database/sql"
	"resto/entities"
)

type TableRestoRepository interface {
	AddTable(table entities.TableResto) error
	GetAllTables() ([]entities.TableResto, error)
}

func (r *Repository) AddTable(table entities.TableResto) error {
	_, err := r.DB.Exec(
		queryInputTable,
		table.Chairs,
		table.IsUsed,
		table.Created_at,
		table.Updated_at,
	)
	return err
}

func (r *Repository) GetAllTables() ([]entities.TableResto, error) {
	var tables []entities.TableResto

	rows, err := r.DB.Query(queryGetAllTables)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		var table = entities.TableResto{}
		var err = rows.Scan(
			&table.Number,
			&table.Chairs,
			&table.IsUsed,
			&table.Created_at,
			&table.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, err
}
