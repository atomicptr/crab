package urllist

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func FetchUrlsFromPath(path string, client *http.Client) ([]string, error) {
	var urls []string

	data, err := fetchList(path, client)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		// url without hostname
		if strings.HasPrefix(line, "/") {
			urls = append(urls, line)
			continue
		}

		// try to parse the url, if no errors are found, append it to the list
		u, err := url.Parse(line)
		if err == nil {
			urls = append(urls, u.String())
		}
	}

	return urls, nil
}

func fetchList(path string, client *http.Client) (io.Reader, error) {
	if strings.HasPrefix(path, "http") {
		return fetchListFromWeb(path, client)
	}
	return fetchListFromPath(path)
}

func fetchListFromWeb(path string, client *http.Client) (io.Reader, error) {
	resp, err := client.Get(path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}

func fetchListFromPath(path string) (io.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
