package apiserver

import (
	"net/http"

	"github.com/golovpeter/http-testapi/internal/store/sqlstore"
	"github.com/jmoiron/sqlx"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)

	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
