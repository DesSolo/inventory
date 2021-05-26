package server

import (
	"encoding/json"
	"inventory/internal/collector"
	"log"
	"net/http"
	"strconv"
)

type clientApi struct {
	rs *RestServer
}

func (c *clientApi) IsExist(w http.ResponseWriter, r *http.Request) {
	wh := r.URL.Query().Get("wh")
	if wh != "" {
		hi, err := c.rs.storage.GetByWH(wh)
		if err != nil {
			log.Printf("fault call storage err: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(
			[]byte(strconv.FormatBool(hi != nil)),
		)
		return

	}

	serial := r.URL.Query().Get("serial")
	if serial != "" {
		hi, err := c.rs.storage.GetBySerial(serial)
		if err != nil {
			log.Printf("fault call storage err: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(
			[]byte(strconv.FormatBool(hi != nil)),
		)
		return

	}
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func (c *clientApi) Upload(checkExist bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hi collector.HostInfo
		if err := json.NewDecoder(r.Body).Decode(&hi); err != nil {
			log.Printf("fault decode json err: %s", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if err := hi.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if checkExist {
			isExist, err := c.rs.storage.IsExist(hi)
			if err != nil {
				log.Printf("fault call storage err: %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			if isExist {
				w.WriteHeader(http.StatusOK)
			}
		}

		if err := c.rs.storage.Save(hi); err != nil {
			log.Printf("fault call storage err: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		log.Printf("%#v", hi)
		w.WriteHeader(http.StatusCreated)
	}
}
