package crawler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
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

		// Write Json File if OutputJson flag is set
		if c.OutputJson != "" {
			c.writeJsonFile(message, c.OutputJson)
		}

		// Write to file if OutputFile flag is set
		if c.OutputFile != "" {
			c.writeLineToFIle(message, c.OutputFile)
		}
		c.printMutex.Unlock()

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
		`{"status": %d, "url": "%s", "time": %d, "duration": %d, "error": %s}`,
		218,
		url,
		time.Now().Unix(),
		duration.Milliseconds(),
		escapeString(err.Error()),
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

// Check File path and create if not exists
func (c *Crawler) checkFileAndCreate(filePath string) {
	// Check directory and create if not exists
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Check file and create if not exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Write files to the output writer
func (c *Crawler) writeLineToFIle(message, filePath string) {
	// filePath check and create
	c.checkFileAndCreate(filePath)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var error string
	var status float64
	var url string
	var time float64
	var duration float64

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		log.Fatal(err)
	}

	if reflect.TypeOf(data["url"]) != nil || reflect.TypeOf(data["time"]) != nil || reflect.TypeOf(data["duration"]) != nil {
		status = data["status"].(float64)
		url = data["url"].(string)
		time = data["time"].(float64)
		duration = data["duration"].(float64)

		if reflect.TypeOf(data["error"]) != nil && data["error"] != "" {
			error = data["error"].(string)
			_, err = file.WriteString(fmt.Sprintf("%d\t%s\t%d\t%d\t%s", int(status), url, int(time), int(duration), error) + "\n")
		} else {
			_, err = file.WriteString(fmt.Sprintf("%d\t%s\t%d\t%d", int(status), url, int(time), int(duration)) + "\n")
		}
	} else {
		_, err = file.WriteString(message + "\n")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Write json files to array output
func (c *Crawler) writeJsonFile(message interface{}, filePath string) {
	// filePath check and create
	c.checkFileAndCreate(filePath)

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// file is empty create a new array
	if len(file) == 0 {
		file = []byte("[]")
	}

	var temp []interface{}
	if err := json.Unmarshal(file, &temp); err != nil {
		log.Fatal(err)
	}

	temp = append(temp, message)

	jsonData, err := json.MarshalIndent(temp, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
