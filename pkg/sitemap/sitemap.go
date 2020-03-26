package sitemap

import (
	"github.com/beevik/etree"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func FetchUrlsFromPath(path string, client *http.Client) ([]string, error) {
	var urls []string

	xmlDataBlob, err := fetchXml(path, client)
	if err != nil {
		return nil, err
	}

	doc := etree.NewDocument()
	if _, err := doc.ReadFrom(xmlDataBlob); err != nil {
		return nil, err
	}

	// check if the sitemap is a sitemap index
	sitemapIndex := doc.FindElement("sitemapindex")

	if sitemapIndex != nil {
		for _, sitemap := range sitemapIndex.ChildElements() {
			loc := sitemap.FindElement("loc")
			if loc != nil {
				sitemapUrls, err := FetchUrlsFromPath(loc.Text(), client)
				if err != nil {
					return nil, err
				}
				urls = append(urls, sitemapUrls...)
			}
		}
	}

	// regular sitemap
	urlSet := doc.FindElement("urlset")

	if urlSet != nil {
		for _, urlElement := range urlSet.ChildElements() {
			loc := urlElement.FindElement("loc")
			if loc != nil {
				url := loc.Text()
				urls = append(urls, url)
			}
		}
	}

	return urls, nil
}

func fetchXml(path string, client *http.Client) (io.Reader, error) {
	if strings.HasPrefix(path, "http") {
		return fetchXmlFromWeb(path, client)
	}
	return fetchXmlFromFile(path)
}

func fetchXmlFromWeb(path string, client *http.Client) (io.Reader, error) {
	resp, err := client.Get(path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	return resp.Body, nil
}

func fetchXmlFromFile(path string) (io.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
