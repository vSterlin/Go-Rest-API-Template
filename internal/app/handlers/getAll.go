package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vSterlin/api-template/internal/app/repository"
)

func GetAllHandler(tr *repository.TemplateRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates, err := tr.FindAll()
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(templates)
	}
}
