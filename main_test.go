package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "strings"
)

func TestGetIndex(t *testing.T) {
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

  var contentTypeExpected string = "text/html; charset=UTF-8"

  if contentType, present := rr.HeaderMap["Content-Type"]; present {
    // convert []string to string 
    contentType := strings.Join(contentType, "")
    if contentType != contentTypeExpected {
      t.Errorf("handler returned wrong content type: got %v want %v", contentType, contentTypeExpected)
    }
  } else {
    t.Errorf("handler returned no Content-Type set")
  }
}

func TestPostIndex(t *testing.T) {
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