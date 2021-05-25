package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type adminApi struct {
	rs *RestServer
}

func (a *adminApi) Uploads() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hil, err := a.rs.storage.GetAll()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		resp, err := json.Marshal(hil)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func (a *adminApi) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serial := chi.URLParam(r, "serial")
		if serial == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if err := a.rs.storage.DeleteBySerial(serial); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
