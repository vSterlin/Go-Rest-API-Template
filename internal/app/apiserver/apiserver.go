package apiserver

import (
	"net/http"

	"github.com/vSterlin/api-template/internal/app/store"

	// Postgres driver
	_ "github.com/lib/pq"
	"github.com/vSterlin/api-template/internal/app/config"
)

// Start starts up the server
func Start(config *config.Config) error {
	db, err := store.New(config.Store)
	if err != nil {
		return nil
	}
	defer db.Close()

	srv := newServer(db)
	return http.ListenAndServe(config.Server.BindAddr, srv)
}
