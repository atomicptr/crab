package crawler

import (
	"fmt"
	"time"
)

func log(statusCode int, url string, duration time.Duration) {
	message := fmt.Sprintf(
		`{"status": %d, "url": "%s", "time": %d, "duration": %d}`,
		statusCode,
		url,
		time.Now().Unix(),
		duration.Milliseconds(),
	)
	fmt.Println(message)
}

func logError(err error, url string, duration time.Duration) {
	message := fmt.Sprintf(
		`{"err": %s, "url": "%s", "time": %d, "duration": %d}\n`,
		err,
		url,
		time.Now().Unix(),
		duration.Milliseconds(),
	)
	fmt.Println(message)
}
