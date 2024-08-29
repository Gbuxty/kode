package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"kode/handlers"
)

func TestCreateNoteHandler(t *testing.T) {
	payload := `{"id":"1", "content":"This is a test note"}`
	req, err := http.NewRequest("POST", "/notes", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateNoteHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetNotesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/notes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetNotesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
