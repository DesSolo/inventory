package server

import (
	"encoding/json"
	"net/http"
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
