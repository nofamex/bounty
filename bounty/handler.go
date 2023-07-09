package bounty

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
)

func BountyHandler() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createBountyHandler)
	return r
}

func createBountyHandler(w http.ResponseWriter, r *http.Request) {
	var body CreateBountyRequest
	json.NewDecoder(r.Body).Decode(&body)

	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := createBounty(r.Context(), body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(res)
}
