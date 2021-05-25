package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"inventory/internal/collector"
	"io"
	"net/http"
)

type Rest struct {
	url    string
	token  string
	client *http.Client
}

func NewRest(url, token string) *Rest {
	return &Rest{
		url:    url,
		token:  token,
		client: &http.Client{},
	}
}

func (r *Rest) newRequest(method, uri string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, r.url+uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Inventory-Token", r.token)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (r *Rest) Upload(inf *collector.HostInfo) error {
	data, err := json.Marshal(inf)
	if err != nil {
		return err
	}

	req, err := r.newRequest("POST", "/upload", bytes.NewReader(data))
	if err != nil {
		return err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("fault call external storage code: %d", resp.StatusCode)
	}

	return nil
}
