package logging

import (
	"encoding/json"
	"fmt"
	"time"
)

// Service unified
type Service interface {
	Debug(string, string)
	Info(string, string)
	Warn(string, string)
	Error(string, error)
}

type defaultLogging struct {
	logLevel string
}

// NewStdoutLogging constructor of the logging service
func NewStdoutLogging(ll string) Service {
	return &defaultLogging{logLevel: ll}
}

// Logs a DEBUG level message
func (dl *defaultLogging) Debug(fn, msg string) {
	if dl.logLevel == "DEBUG" {
		dl.print("DEBUG", fn, msg, "")
	}
}

// Logs an INFO level message
func (dl *defaultLogging) Info(fn, msg string) {
	if dl.logLevel == "DEBUG" || dl.logLevel == "INFO" {
		dl.print("INFO", fn, msg, "")
	}
}

// Logs a WARN level message
func (dl *defaultLogging) Warn(fn, msg string) {
	if dl.logLevel == "DEBUG" || dl.logLevel == "INFO" || dl.logLevel == "WARN" {
		dl.print("WARN", fn, msg, "")
	}
}

// Logs an ERROR level message
func (dl *defaultLogging) Error(fn string, err error) {
	if dl.logLevel == "DEBUG" || dl.logLevel == "INFO" || dl.logLevel == "WARN" || dl.logLevel == "ERROR" {
		dl.print("ERROR", fn, "", err.Error())
	}
}

// print prints logging message to console
func (dl *defaultLogging) print(logLevel, fn, msg, errStr string) {
	l := Log{
		TimeStamp: time.Now(),
		Message:   msg,
		Exception: errStr,
		LogLevel:  logLevel,
	}

	pp, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(pp))
}
