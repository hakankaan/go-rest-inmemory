package logging

import "time"

type Log struct {
	TimeStamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Exception string    `json:"exception"`
	LogLevel  string    `json:"logLevel"`
}
