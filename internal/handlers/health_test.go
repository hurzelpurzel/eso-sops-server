package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	r := httptest.NewRecorder()

	HealthHandler(r, req)

	if r.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", r.Code)
	}

	if ct := r.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected content-type application/json, got %s", ct)
	}
}
