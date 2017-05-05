package main

import (
  "encoding/json"
  "fmt"
  "crypto/rand"
  "encoding/hex"
  "errors"
)

func BuildData(body []byte, event Event, sessionId string) error {
  switch event.Type {
    case "resize":
      if err := BuildResizeData(body, event, sessionId); err != nil {
        return err
      }
    case "copyAndPaste":
      if err := BuildCopyAndPasteData(body, event, sessionId); err != nil {
        return err
      }
    case "timeTaken":
      if err := BuildFormCompletionTimeData(body, event, sessionId); err != nil {
        return err
      }
    default:
      return errors.New("Error")
  }
  PrettyPrint(sessions[sessionId])
  return nil  
}

func BuildResizeData(body []byte, event Event, sessionId string) error {
  var resizeData ResizeData
  if err := json.Unmarshal(body, &resizeData); err != nil {
    return err
  }
  sessions[sessionId].ResizeFrom.Width  = resizeData.ResizeFromWidth
  sessions[sessionId].ResizeFrom.Height = resizeData.ResizeFromHeight
  sessions[sessionId].ResizeTo.Width    = resizeData.ResizeToWidth
  sessions[sessionId].ResizeTo.Height   = resizeData.ResizeToHeight
  return nil
}

func BuildCopyAndPasteData(body []byte, event Event, sessionId string) error {
  var copyAndPasteData CopyAndPasteData
  if err := json.Unmarshal(body, &copyAndPasteData); err != nil {
    return err
  }
  fieldId := copyAndPasteData.FieldId
  pasted  := copyAndPasteData.Pasted
  sessions[sessionId].CopyAndPaste[fieldId] = pasted
  return nil
}

func BuildFormCompletionTimeData(body []byte, event Event, sessionId string) error {
  var formCompletionTimeData FormCompletionTimeData
  if err := json.Unmarshal(body, &formCompletionTimeData); err != nil {
    return err
  }
  time := formCompletionTimeData.FormCompletionTime
  sessions[sessionId].FormCompletionTime = time
  return nil
}

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

