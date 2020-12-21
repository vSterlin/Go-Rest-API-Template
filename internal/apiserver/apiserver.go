package apiserver

import (
	"net/http"

	"github.com/vSterlin/api-template/internal/config"
	"github.com/vSterlin/api-template/internal/store"
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
