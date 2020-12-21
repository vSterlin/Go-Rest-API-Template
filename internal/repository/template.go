package repository

import (
	"database/sql"

	"github.com/vSterlin/api-template/internal/model"
)

//go:generate go run github.com/golang/mock/mockgen -destination mocks/store_mock.go -package mocks github.com/vSterlin/api-template/internal/repository TemplateStore
type TemplateStore interface {
	Query(q string, args ...interface{}) (*sql.Rows, error)
	QueryRow(q string, args ...interface{}) *sql.Row
}

type TemplateRepo struct {
	s TemplateStore
}

func NewTemplate(s TemplateStore) *TemplateRepo {
	return &TemplateRepo{
		s: s,
	}
}

func (tr *TemplateRepo) FindAll() ([]*model.Template, error) {

	rows, err := tr.s.Query("SELECT * FROM template;")
	if err != nil {
		return nil, err
	}

	pmArr := []*model.Template{}
	for rows.Next() {
		pm := &model.Template{}
		rows.Scan(&pm.ID, &pm.Name)
		pmArr = append(pmArr, pm)
	}
	return pmArr, nil

}
func (tr *TemplateRepo) FindById(id int) (*model.Template, error) {
	row := tr.s.QueryRow("SELECT * FROM template WHERE id=$1;", id)
	pm := &model.Template{}

	err := row.Scan(&pm.ID, &pm.Name)
	if err != nil {
		return nil, err
	}
	return pm, nil
}
