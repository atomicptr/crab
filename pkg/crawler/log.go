package crawler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
		defer c.printMutex.Unlock()

		logsToFile := false

		if c.OutputJson != "" {
			c.writeLineToJsonFile(message, c.OutputJson)
			logsToFile = true
		}

		if c.OutputFile != "" {
			c.writeLineToFile(message, c.OutputFile)
			logsToFile = true
		}

		// dont log to stdout if we log to file
		if logsToFile {
			return
		}

		_, _ = fmt.Fprintln(c.OutWriter, message)
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

// assureFileExists Check if file path exists and create the file otherwise
func (c *Crawler) assureFileExists(filePath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}

// writeLineToFile Writes message to file
func (c *Crawler) writeLineToFile(message, filePath string) {
	err := c.assureFileExists(filePath)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	data := struct {
		Err      string
		Status   int
		Url      string
		Time     int
		Duration int
	}{}
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		log.Fatal(err)
	}

	if data.Err != "" {
		_, err = file.WriteString(fmt.Sprintf("%s\t%s\t%d\t%dms", data.Err, data.Url, data.Time, data.Duration) + "\n")
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	_, err = file.WriteString(fmt.Sprintf("%d\t%s\t%d\t%dms", data.Status, data.Url, data.Time, data.Duration) + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

// writeJsonFile Write message to json file
func (c *Crawler) writeLineToJsonFile(message, filePath string) {
	err := c.assureFileExists(filePath)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	if len(file) == 0 {
		file = []byte("[]")
	}

	var temp []map[string]interface{}
	if err := json.Unmarshal(file, &temp); err != nil {
		log.Fatal(err)
	}

	var msg map[string]interface{}
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		log.Fatal(err)
	}

	temp = append(temp, msg)

	jsonData, err := json.MarshalIndent(temp, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
