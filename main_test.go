package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestGetIndex(t *testing.T) {
  // Create a request to pass to our handler.
  req, err := http.NewRequest("GET", "/", nil)
  if err != nil {
    t.Fatal(err)
  }
  
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(GetIndex)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }
}

func TestPostIndex(t *testing.T) {
  // Create a request to pass to our handler.
  req, err := http.NewRequest("POST", "/", nil)
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(GetIndex)

  handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusMethodNotAllowed {
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
  }
}