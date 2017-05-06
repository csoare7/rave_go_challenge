package main

import (
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
  "io"
)

const allowedOrigin string = "http://localhost:8000" 

var sessions = make(map[string] *Data)

func PostData(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
      http.Error(w, "Bad request", http.StatusBadRequest)
      return
    }

    var baseData BaseData
    if err := json.Unmarshal(body, &baseData); err != nil {
      http.Error(w, "Bad request", http.StatusBadRequest)
      return
    }  

    sessionId, err := ReadOrInitSessionId(&baseData); if err != nil {
      http.Error(w, "Internal server error", http.StatusInternalServerError)
      return
    }

    var event Event
    if err := json.Unmarshal(body, &event); err != nil {
      http.Error(w, "Bad request", http.StatusBadRequest)
      return
    }

    if err := BuildData(body, event, sessionId); err != nil {
      log.Println("1")
      http.Error(w, "Bad request", http.StatusBadRequest)
      return
    }

    var responseObj = make(map[string]string) 
    responseObj["sessionId"] = sessionId
    responseJson, err := json.Marshal(responseObj)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    setAccessControlHeaders(w)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    w.Write(responseJson)
  } else if r.Method == "OPTIONS" {
    setAccessControlHeaders(w)
    return
  } else {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    return
  }
}

func setAccessControlHeaders(w http.ResponseWriter) {
  w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
  w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding")
}

func main() {
  router := http.NewServeMux()
  router.HandleFunc("/data", PostData)

  err := http.ListenAndServe(":3000", router)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
