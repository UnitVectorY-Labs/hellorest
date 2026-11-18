package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	w := httptest.NewRecorder()
	helloWorld(w, httptest.NewRequest(http.MethodGet, "/", nil))

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	if got := w.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", got)
	}
	if got := w.Header().Get("X-App-Version"); got != Version {
		t.Errorf("X-App-Version = %q, want %q", got, Version)
	}
	if got := w.Body.String(); got != "{\"hello\":\"world\"}\n" {
		t.Errorf("body = %q, want JSON response followed by a newline", got)
	}
}

func TestHelloWorldNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	helloWorld(w, httptest.NewRequest(http.MethodGet, "/missing", nil))
	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", w.Code, http.StatusNotFound)
	}
}
