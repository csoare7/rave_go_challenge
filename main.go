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

var sessions = make(map[string]*Data) 

func ReadOrInitSessionId(bD *BaseData) (string, error) {
  // provide simple random generator for session
  // https://astaxie.gitbooks.io/build-web-application-with-golang/en/06.2.html
  if bD.SessionId == "" {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
      return "", err
    }
    return hex.EncodeToString(bytes), nil
  }
  return bD.SessionId, nil
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

    fmt.Println(sessionId)

    var event Event
    if err := json.Unmarshal(body, &event); err != nil {
      http.Error(w, "Bad request", http.StatusBadRequest)
      return
    }



    // switch event
    //   case "resize":
    //   case "copyAndPaste":

    //   case "timeTaken":
    //       sessions[sessionId].FormCompletionTime = time
    //   default:
    //     http.Error(w, "Bad request", http.StatusBadRequest)
    //     return


    var time FormCompletionTimeData
    if err := json.Unmarshal(body, &time); err != nil {
      http.Error(w, "Bad request", http.StatusBadRequest)
      return
    }


    // var data Data//map[string] json.RawMessage
    // if err := json.Unmarshal(body, &data); err != nil {
    //   http.Error(w, "Bad request", http.StatusBadRequest)
    //   return
    // }
    fmt.Println(time)

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
