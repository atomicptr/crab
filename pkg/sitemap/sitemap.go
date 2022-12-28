package sitemap

import (
	"github.com/atomicptr/crab/pkg/crawler"
	"github.com/beevik/etree"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func FetchUrlsFromPath(path string, client *http.Client, modifier *crawler.RequestModifier) ([]string, error) {
	var urls []string

	xmlDataBlob, err := fetchXml(path, client, modifier)
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
				sitemapUrls, err := FetchUrlsFromPath(loc.Text(), client, modifier)
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

func fetchXml(path string, client *http.Client, modifier *crawler.RequestModifier) (io.Reader, error) {
	if strings.HasPrefix(path, "http") {
		return fetchXmlFromWeb(path, client, modifier)
	}
	return fetchXmlFromFile(path)
}

func fetchXmlFromWeb(path string, client *http.Client, modifier *crawler.RequestModifier) (io.Reader, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	if modifier != nil {
		modifier.Do(req)
	}

	resp, err := client.Do(req)
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
