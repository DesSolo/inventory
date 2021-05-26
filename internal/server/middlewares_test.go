package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getEmptyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func TestDefaultContentType(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	mw := DefaultContentType("application/json")(getEmptyHandler())
	mw.ServeHTTP(w, r)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestTokenAuthMiddlewareEmpty(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	mw := TokenAuthMiddleware("secret")(getEmptyHandler())
	mw.ServeHTTP(w, r)
	assert.Equal(t, http.StatusForbidden, w.Result().StatusCode)
}

func TestTokenAuthMiddlewareValid(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set("Inventory-Token", "secret")
	w := httptest.NewRecorder()

	mw := TokenAuthMiddleware("secret")(getEmptyHandler())
	mw.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
