package main

import (
  "net/http"
  "testing"
  "strings"
  "fmt"
)

func TestPostData(t *testing.T) {

  timeJson := `'{"eventType": "timeTaken", "websiteUrl": "https://ravelin.com", "sessionId": "", "time": 100}'`
  reader := strings.NewReader(timeJson) //Convert string to reader
  
  r, err := http.NewRequest("POST", "http://localhost:3000/data", reader)
  r.Header.Add("Content-Type", "application/json")
  response, err := http.DefaultClient.Do(r)

  if err != nil {
      t.Error(err)
  }
  body := response.Body

  fmt.Println(body)



}