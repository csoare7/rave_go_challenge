package main

import (
  "net/http"
  "log"
  //"encoding/json"
  //"io/ioutil"
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
    var data interface{}
    err := json.Unmarshal(b, &data)

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
