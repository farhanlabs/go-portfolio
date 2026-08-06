package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(homeHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d but got %d", http.StatusOK, rr.Code)
	}

	expected := "text/html; charset=utf-8"
	if rr.Header().Get("Content-Type") != expected {
		t.Errorf("expected Content-Type %s but got %s",
			expected,
			rr.Header().Get("Content-Type"))
	}
}