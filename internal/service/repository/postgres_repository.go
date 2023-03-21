package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (pr *PostgresRepository) Set(id uint64, url string) error {
	_, err := pr.Get(id)
	if err == nil {
		return ErrorAlreadyExists
	}

	_, err = pr.db.Exec("INSERT INTO urls (id, url) VALUES($1, $2)", id, url)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostgresRepository) Get(id uint64) (string, error) {
	var result string

	err := pr.db.Get(&result, "SELECT url FROM urls WHERE id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrorNotFound
		}
		return "", err
	}

	return result, nil
}

func (pr *PostgresRepository) GetLastIndex() uint64 {
	var result uint64

	err := pr.db.Get(&result, "SELECT MAX(id) FROM urls")
	if err != nil {
		return 0
	}

	return result
}
