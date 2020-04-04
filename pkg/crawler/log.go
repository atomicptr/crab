package crawler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/atomicptr/crab/pkg/filter"
)

// safePrintln logs a message protected by a mutex
func (c *Crawler) safePrintln(statusCode int, message string) {
	if c.statusFilter == nil {
		c.statusFilter = filter.NewFilter()
	}

	if c.statusFilter.IsValid(c.FilterStatusQuery, int64(statusCode)) {
		c.printMutex.Lock()
		_, _ = fmt.Fprintln(c.OutWriter, message)
		c.printMutex.Unlock()
	}
}

// log logs a json log with the status code, url, timestamp and duration of the request
func (c *Crawler) log(statusCode int, url string, duration time.Duration) {
	message := fmt.Sprintf(
		`{"status": %d, "url": "%s", "time": %d, "duration": %d}`,
		statusCode,
		url,
		time.Now().Unix(),
		duration.Milliseconds(),
	)
	c.safePrintln(statusCode, message)
}

// logError logs a json log with an error, url, timestamp and duration of the request
func (c *Crawler) logError(err error, url string, duration time.Duration) {
	message := fmt.Sprintf(
		`{"err": %s, "url": "%s", "time": %d, "duration": %d}`,
		escapeString(err.Error()),
		url,
		time.Now().Unix(),
		duration.Milliseconds(),
	)
	c.safePrintln(218, message)
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
