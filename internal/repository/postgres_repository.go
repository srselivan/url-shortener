package repository

import "github.com/jmoiron/sqlx"

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (pr *PostgresRepository) Set(key string, url string) error {
	_, err := pr.db.Exec("INSERT INTO urls (key, url) VALUES($1, $2)", key, url)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostgresRepository) Get(key string) (string, error) {
	var result string

	err := pr.db.Get(&result, "SELECT url FROM urls WHERE key = $1", key)
	if err != nil {
		return "", err
	}

	return result, nil
}
