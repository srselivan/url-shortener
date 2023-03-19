package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func New(url string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		return nil, err
	}

	return db, err
}
