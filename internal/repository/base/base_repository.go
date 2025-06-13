package base

import "database/sql"

type BaseRepository[T any] interface {
	GetAll() ([]T, error)
	GetByID(id int64) (*T, error)
	Delete(id int64) error
}

type BaseSQLRepository[T any] struct {
	DB        *sql.DB
	TableName string
	ScanFunc  func(*sql.Rows) (T, error)
}

func (r *BaseSQLRepository[T]) GetAll() ([]T, error) {
	rows, err := r.DB.Query("SELECT * FROM " + r.TableName + " ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []T
	for rows.Next() {
		item, err := r.ScanFunc(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *BaseSQLRepository[T]) GetByID(id int64) (*T, error) {
	query := "SELECT * FROM " + r.TableName + " WHERE id = $1"
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		item, err := r.ScanFunc(rows)
		if err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, sql.ErrNoRows
}

func (r *BaseSQLRepository[T]) Delete(id int64) error {
	query := "DELETE FROM " + r.TableName + " WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
