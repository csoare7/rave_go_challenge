package main

import (
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
  "io"
  "fmt"
)

var sessions = make(map[string] *Data)

func GetIndex(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    w.Header().Set("Content-Type", "text/html; charset=UTF-8")
    http.ServeFile(w, r, "static/index.html")
  } else {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    return
  }
}

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

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusAccepted)
    w.Write(responseJson)
  } else {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    return
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
  fmt.Println("ListenAndServe on port 3000")
}
