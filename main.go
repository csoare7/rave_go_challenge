package main

import (
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
  "io"
  "fmt"
  "crypto/rand"
  "encoding/hex"
)

var sessions = make(map[string] *Data)

// provide simple random generator for hex creation
func ReadOrInitSessionId(bD *BaseData) (string, error) {
  _, ok := sessions[bD.SessionId]; if !ok {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
      return "", err
    }
    sessionId := hex.EncodeToString(bytes)
    sessions[sessionId] = &Data{SessionId: sessionId, CopyAndPaste: make(map[string]bool)}
    return sessionId, nil
  }
  return bD.SessionId, nil
}

// https://siongui.github.io/2016/01/30/go-pretty-print-variable/
func PrettyPrint(data interface{}) {
  d, _ := json.MarshalIndent(data, "", "  ")
  fmt.Printf("%+v", string(d))
}

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

    switch event.Type {
      case "resize":
        var resizeData ResizeData
        if err := json.Unmarshal(body, &resizeData); err != nil {
          fmt.Println(err)
          http.Error(w, "Bad request", http.StatusBadRequest)
          return
        }
        sessions[sessionId].ResizeFrom.Width  = resizeData.ResizeFromWidth
        sessions[sessionId].ResizeFrom.Height = resizeData.ResizeFromHeight
        sessions[sessionId].ResizeTo.Width    = resizeData.ResizeToWidth
        sessions[sessionId].ResizeTo.Height   = resizeData.ResizeToHeight
        PrettyPrint(sessions[sessionId])
      case "copyAndPaste":
        var copyAndPasteData CopyAndPasteData
        if err := json.Unmarshal(body, &copyAndPasteData); err != nil {
          http.Error(w, "Bad request", http.StatusBadRequest)
          return
        }
        fieldId := copyAndPasteData.FieldId
        pasted  := copyAndPasteData.Pasted
        sessions[sessionId].CopyAndPaste[fieldId] = pasted
        PrettyPrint(sessions[sessionId])
      case "timeTaken":
        var formCompletionTimeData FormCompletionTimeData
        if err := json.Unmarshal(body, &formCompletionTimeData); err != nil {
          http.Error(w, "Bad request", http.StatusBadRequest)
          return
        }
        time := formCompletionTimeData.FormCompletionTime
        sessions[sessionId].FormCompletionTime = time
        PrettyPrint(sessions[sessionId])
      default:
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusAccepted)

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
}
