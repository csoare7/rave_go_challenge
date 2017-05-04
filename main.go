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

type Data struct {
  WebsiteUrl         string `json:"websiteUrl"`
  SessionId          string `json:"sessionId"`
  ResizeFrom         Dimension `json:"resizeFrom"`
  ResizeTo           Dimension `json:"resizeTo"`
  CopyAndPaste       map[string]bool `json:"width"`// map[fieldId]true
  FormCompletionTime int `json:"time"`// Seconds
}

type Dimension struct {
  Width  string `json:"width"`
  Height string `json:"height"`
}

type BaseData struct {
  WebsiteUrl         string `json:"websiteUrl"`
  SessionId          string `json:"sessionId"` 
}

type ResizeData struct {
  *BaseData
  ResizeFrom         Dimension `json:"resizeFrom"`
  ResizeTo           Dimension `json:"resizeTo"`
}

type CopyAndPasteData struct {
  *BaseData
  CopyAndPaste       map[string]bool
}

type FormCompletionTimeData struct {
  *BaseData
  FormCompletionTime int `json:"time"`
}

type Event struct {
  Type string `json:"eventType"`
}

var sessions = make(map[string] *Data)

// provide simple random generator for hex creation
func ReadOrInitSessionId(bD *BaseData) (string, error) {
  _, ok := sessions[bD.SessionId]; if !ok {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
      return "", err
    }
    sessionId := hex.EncodeToString(bytes)
    sessions[sessionId] = &Data{SessionId: sessionId}
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
    fmt.Println(sessionId)
    fmt.Println(event.Type)
    switch event.Type {
      // case "resize":
      //   var resize ResizeData
      //   if err := json.Unmarshal(body, &resize); err != nil {
      //     http.Error(w, "Bad request", http.StatusBadRequest)
      //     return
      //   }
      // case "copyAndPaste":
      //   var copyAndPaste CopyAndPasteData
      //   if err := json.Unmarshal(body, &copyAndPasteData); err != nil {
      //     http.Error(w, "Bad request", http.StatusBadRequest)
      //     return
      //   }
      case "timeTaken":
        var formCompletionTimeData FormCompletionTimeData
        if err := json.Unmarshal(body, &formCompletionTimeData); err != nil {
          http.Error(w, "Bad request", http.StatusBadRequest)
          return
        }
        sessions[sessionId].FormCompletionTime = formCompletionTimeData.FormCompletionTime
        PrettyPrint(sessions[sessionId])

      default:
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }



    // var data Data//map[string] json.RawMessage
    // if err := json.Unmarshal(body, &data); err != nil {
    //   http.Error(w, "Bad request", http.StatusBadRequest)
    //   return
    // }
    // fmt.Println(time)

    // sessionId, okSessionId := data["sessionId"]; 
    // eventType, okEventType := data["eventType"]; 
    // websiteUrl, okWebsiteUrl := data["websiteUrl"]; if !okSessionId || !okEventType || !okWebsiteUrl {
    //   http.Error(w, "Bad request", http.StatusBadRequest)
    //   return
    // } else {
    //   if _, ok := sessionId.(string); !ok {
    //     http.Error(w, "Bad request", http.StatusBadRequest)
    //     return
    //   }
    //   switch eventType {
    //     case "resize":
    //       // var resizeFrom Dimension := {data["resizeFrom"]["width"], data["resizeFrom"]["height"]}
    //       // var resizeTo Dimension := {data["resizeTo"]["width"], data["resizeTo"]["height"]}

    //       resizeFromWidth, ok := data["resizeFromWidth"].(string); 
    //       resizeFromHeight, ok := data["resizeFromHeight"].(string); if ok {
    //         resizeFrom := new(Dimension).SetValues(resizeFromWidth, resizeFromHeight)
    //         fmt.Println(resizeFrom)
    //       }
    //       // fmt.Printf("%+v", sessions[sessionId])
    //       // resizeTo := new(Dimension).SetValues(data["resizeToWidth"], data["resizeToHeight"])
    //       // sessions[sessionId].resizeFrom = resizeFrom
    //       // sessions[sessionId].resizeTo = resizeTo
    //     case "copyAndPaste":
    //       fmt.Printf("%s", data["pasted"])
    //     case "timeTaken":
    //       fmt.Println(websiteUrl)
    //       if time, ok := data["time"].(int); ok {
    //         sessions[sessionId].FormCompletionTime = time
    //       }
    //     default:
    //       http.Error(w, "Bad request", http.StatusBadRequest)
    //       return
    //   }
    // }

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
