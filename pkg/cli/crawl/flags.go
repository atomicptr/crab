package crawl

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type crawlerFlagOptions struct {
	NumberOfWorkers   int
	PrefixUrl         string
	HttpTimeout       time.Duration
	AuthUsername      string
	AuthPassword      string
	CookieStrings     []string
	HeaderStrings     []string
	FilterStatusQuery string
	cookieMap         map[string]string
	headerMap         map[string]string
	OutputFile        string
	OutputJson        string
}

const (
	defaultNumberOfWorkers = 4
	defaultHttpTimeout     = 30 * time.Second
	defaultOutputFile      = "./output/output.txt"
)

func newCrawlerFlagOptions() crawlerFlagOptions {
	return crawlerFlagOptions{
		NumberOfWorkers: defaultNumberOfWorkers,
		HttpTimeout:     defaultHttpTimeout,
		OutputFile:      defaultOutputFile,
	}
}

func validateCrawlerFlagOptions(flagOptions crawlerFlagOptions) error {
	if flagOptions.NumberOfWorkers <= 0 {
		return fmt.Errorf("number of workers should be at least 1")
	}

	if flagOptions.PrefixUrl != "" && !strings.HasPrefix(flagOptions.PrefixUrl, "http") {
		return fmt.Errorf("prefix url is not a proper url: %s", flagOptions.PrefixUrl)
	}

	if err := validateKeyValueStrings("cookie", flagOptions.CookieStrings); err != nil {
		return err
	}

	if err := validateKeyValueStrings("header", flagOptions.HeaderStrings); err != nil {
		return err
	}

	return nil
}

func validateKeyValueStrings(name string, keyValueSet []string) error {
	regex, err := regexp.Compile(`.+=.+`)
	if err != nil {
		return err
	}

	for _, keyValuePair := range keyValueSet {
		if !regex.MatchString(keyValuePair) {
			return fmt.Errorf(
				"%s does not match pattern %s_name=%s_value for: %s",
				name,
				name,
				name,
				keyValuePair,
			)
		}
	}
	return nil
}

func (flagOptions *crawlerFlagOptions) HeaderMap() map[string]string {
	if flagOptions.headerMap == nil {
		flagOptions.headerMap = createMapFromKeyValueStrings(flagOptions.HeaderStrings)
	}
	return flagOptions.headerMap
}

func (flagOptions *crawlerFlagOptions) CookieMap() map[string]string {
	if flagOptions.cookieMap == nil {
		flagOptions.cookieMap = createMapFromKeyValueStrings(flagOptions.CookieStrings)
	}
	return flagOptions.cookieMap
}

func createMapFromKeyValueStrings(kvStrings []string) map[string]string {
	newMap := make(map[string]string)
	for _, kv := range kvStrings {
		parts := strings.Split(kv, "=")
		newMap[parts[0]] = parts[1]
	}
	return newMap
}
