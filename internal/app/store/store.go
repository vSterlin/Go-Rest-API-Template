package store

import (
	"database/sql"
)

type Store struct {
	db                 *sql.DB
	TemplateRepository *TemplateRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Template() *TemplateRepository {
	if s.TemplateRepository != nil {
		return s.TemplateRepository
	}

	s.TemplateRepository = &TemplateRepository{
		store: s,
	}

	return s.TemplateRepository
}
