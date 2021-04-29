package storage

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"inventory/internal/collector"
	"net/http"
)

type ExternalRest struct {
	address  string
	username string
	password string
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))

}

func (r *ExternalRest) Send(inf *collector.HostInfo) error {
	data, err := json.Marshal(inf)
	if err != nil {
		return err
	}
	client := &http.Client{}

	req, err := http.NewRequest("POST", r.address, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(r.username, r.password))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fault call external storage code: %d", resp.StatusCode)
	}

	return nil
}

func NewRest(address, username, password string) *ExternalRest {
	return &ExternalRest{
		address,
		username,
		password,
	}
}
