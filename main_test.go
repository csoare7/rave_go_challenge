package main

import (
  "net/http"
  "testing"
  "strings"
  "encoding/json"
)

func TestPostData(t *testing.T) {

  timeJson := `{"eventType": "timeTaken", "websiteUrl": "https://ravelin.com", "sessionId": "", "time": 100}`
  reader := strings.NewReader(timeJson) //Convert string to reader
  
  r, err := http.NewRequest("POST", "http://localhost:3000/data", reader)
  r.Header.Add("Content-Type", "application/json")
  response, err := http.DefaultClient.Do(r)

  if err != nil {
    t.Error(err)
  }

  result := make(map[string]string)
  if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
    t.Error(err)
  }

  if sessionId, ok := result["sessionId"]; !ok || sessionId == "" {
    t.Error(err) 
  }

}