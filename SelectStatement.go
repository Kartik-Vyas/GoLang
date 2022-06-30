package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{} // passing an empty struct to close the application
}

func logger() {
	for {
		select { // just like switch statement it will exceute
		case entry := <-logCh:
			fmt.Println(entry.time, entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
