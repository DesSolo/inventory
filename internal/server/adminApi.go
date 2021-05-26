package server

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"inventory/internal/collector"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

type adminApi struct {
	rs *RestServer
}

func (a *adminApi) Uploads() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hil, err := a.rs.storage.GetAll()
		if err != nil {
			log.Printf("fault call storage err: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		resp, err := json.Marshal(hil)
		if err != nil {
			log.Printf("fault marshall json err: %s", err)
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
			log.Printf("fault call storage err: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (a *adminApi) Update() http.HandlerFunc {
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

		if err := a.rs.storage.Save(hi); err != nil {
			log.Printf("fault call storage err: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (a *adminApi) Export() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hil, err := a.rs.storage.GetAll()
		if err != nil {
			log.Printf("fault call storage err: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var b bytes.Buffer
		writer := csv.NewWriter(&b)
		writer.Write([]string{
			"WH", "UserName", "HostName", "SerialNumber", "Manufacturer", "SystemVersion", "MacAddress",
		})
		for _, hi := range hil {
			err := writer.Write([]string{
				hi.WH, hi.UserName, hi.HostName, hi.SerialNumber, hi.Manufacturer, hi.SystemVersion, strings.Join(hi.MacAddress, "|"),
			})
			if err != nil {
				log.Printf("fault write csv err: %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
		writer.Flush()

		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment;filename=InventoryExport.csv")
		w.Write(b.Bytes())
	}
}
