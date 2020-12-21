package store

import (
	"database/sql"
	// Postgres driver
	_ "github.com/lib/pq"
	"github.com/vSterlin/api-template/internal/repository"

	"github.com/vSterlin/api-template/internal/config"
)

type Store struct {
	db           *sql.DB
	templateRepo *repository.TemplateRepo
}

func New(c config.Store) (*Store, error) {

	db, err := sql.Open(c.Driver, c.DSN)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Close() {

	s.db.Close()
}

func (s *Store) Template() *repository.TemplateRepo {
	if s.templateRepo != nil {
		return s.templateRepo
	}

	s.templateRepo = repository.NewTemplate(s.db)

	return s.templateRepo
}
