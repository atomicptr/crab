package crawler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var printMutex sync.Mutex

// safePrintln logs a message protected by a mutex
func safePrintln(message string) {
	printMutex.Lock()
	fmt.Println(message)
	printMutex.Unlock()
}

// log logs a json log with the status code, url, timestamp and duration of the request
func log(statusCode int, url string, duration time.Duration) {
	message := fmt.Sprintf(
		`{"status": %d, "url": "%s", "time": %d, "duration": %d}`,
		statusCode,
		url,
		time.Now().Unix(),
		duration.Milliseconds(),
	)
	safePrintln(message)
}

// logError logs a json log with an error, url, timestamp and duration of the request
func logError(err error, url string, duration time.Duration) {
	message := fmt.Sprintf(
		`{"err": %s, "url": "%s", "time": %d, "duration": %d}`,
		escapeString(err.Error()),
		url,
		time.Now().Unix(),
		duration.Milliseconds(),
	)
	safePrintln(message)
}

// escapeString escapes a string to be used as a json value
func escapeString(str string) string {
	b, err := json.Marshal(str)
	if err != nil {
		// could not parse it, base64 encode it and send that into the log instead
		bStr := base64.StdEncoding.EncodeToString([]byte(str))
		return "base64:" + bStr
	}
	return string(b)
}
