package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newHandler()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Fatalf("handler returned unexpected content type: got %v want %v", contentType, expected)
	}

	if !strings.Contains(rr.Body.String(), "DevOps Engineer") {
		t.Fatalf("expected response body to contain portfolio content")
	}
}

func TestStylesPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/styles.css", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := newHandler()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "text/css; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Fatalf("handler returned unexpected content type: got %v want %v", contentType, expected)
	}
}
