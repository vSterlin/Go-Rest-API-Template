package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vSterlin/api-template/internal/repository"

	"github.com/gorilla/mux"
)

func GetOneHandler(tr *repository.TemplateRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			panic(err)
		}

		template, err := tr.FindById(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		json.NewEncoder(w).Encode(template)
	}
}
