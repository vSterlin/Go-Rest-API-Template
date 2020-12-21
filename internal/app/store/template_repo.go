package store

import "github.com/vSterlin/api-template/internal/app/model"

type TemplateRepository struct {
	store *Store
}

func (pr *TemplateRepository) FindAll() ([]*model.Template, error) {

	rows, err := pr.store.db.Query("SELECT * FROM template;")
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
func (pr *TemplateRepository) FindById(id int) (*model.Template, error) {
	row := pr.store.db.QueryRow("SELECT * FROM template WHERE id=$1;", id)
	pm := &model.Template{}

	err := row.Scan(&pm.ID, &pm.Name)
	if err != nil {
		return nil, err
	}
	return pm, nil
}
