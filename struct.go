package main

type Data struct {
  WebsiteUrl         string `json:"websiteUrl"`
  SessionId          string `json:"sessionId"`
  ResizeFrom         Dimension `json:"resizeFrom"`
  ResizeTo           Dimension `json:"resizeTo"`
  CopyAndPaste       map[string]bool `json:"copyAndPaste"` // map[fieldId]true
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
  ResizeFromWidth    string `json:"resizeFromWidth"`
  ResizeFromHeight   string `json:"resizeFromHeight"`
  ResizeToWidth      string `json:"resizeToWidth"`
  ResizeToHeight     string `json:"resizeToHeight"`
}

type CopyAndPasteData struct {
  *BaseData
  FieldId            string `json:"fieldId"`
  Pasted             bool   `json:"pasted"`
}

type FormCompletionTimeData struct {
  *BaseData
  FormCompletionTime int `json:"time"`
}

type Event struct {
  Type               string `json:"eventType"`
}