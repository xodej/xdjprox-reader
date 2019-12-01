package main

// Entry log entry
type Entry struct {
	Level     string `json:"level"`
	Message   string `json:"msg"`
	RequestID string `json:"request_id"`
	Session   string `json:"session"`
	Time      string `json:"time"`
	Type      string `json:"type"`
}
