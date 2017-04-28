package main

import (
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
  "io"
  "fmt"
)

type Data struct {
  WebsiteUrl         string
  SessionId          string
  ResizeFrom         Dimension
  ResizeTo           Dimension
  CopyAndPaste       map[string]bool // map[fieldId]true
  FormCompletionTime int // Seconds
}

type Dimension struct {
  Width  string
  Height string
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    w.Header().Set("Content-Type", "text/html; charset=UTF-8")
    http.ServeFile(w, r, "static/index.html")
  } else {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
  }
}

func PostData(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
      http.Error(w, "Bad request", http.StatusBadRequest)
    }

    var data map[string]interface{}
    if err := json.Unmarshal(body, &data); err != nil {
      http.Error(w, "Bad request", http.StatusBadRequest)
    }
    
    fmt.Printf("%+v", data)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)

  } else {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
  }
}

func main() {
  router := http.NewServeMux()
  router.HandleFunc("/", GetIndex)
  router.HandleFunc("/data", PostData)

  err := http.ListenAndServe(":3000", router)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
