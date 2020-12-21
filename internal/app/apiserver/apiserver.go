package apiserver

import (
	"database/sql"
	"net/http"

	// Postgres driver
	_ "github.com/lib/pq"
	"github.com/vSterlin/api-template/internal/app/store"
)

// Start starts up the server
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return nil
	}
	defer db.Close()
	store := store.New(db)

	srv := newServer(store)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
