package api

import (
	"time"
)

//TimeEvent represents a time event such as a meeting
type TimeEvent struct {
	Day        time.Time `json:"day"`
	Type       string    `json:"type"`
	Who        string    `json:"who"`
	Duration   uint      `json:"duration"`
	Subject    string    `json:"subject"`
	Department string    `json:"department"`
	Recurrent  bool      `json:"recurrent"`
}
