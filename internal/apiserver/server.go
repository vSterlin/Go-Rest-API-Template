package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/vSterlin/api-template/internal/app/handlers"
	"github.com/vSterlin/api-template/internal/app/store"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  *store.Store
}

func newServer(store *store.Store) *server {

	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	tr := s.store.Template()
	s.router.HandleFunc("/template", handlers.GetAllHandler(tr)).Methods("GET")
	s.router.HandleFunc("/template/{id}", handlers.GetOneHandler(tr)).Methods("GET")

}
