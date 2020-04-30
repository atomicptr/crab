package tools

import (
	"fmt"
	"github.com/atomicptr/crab/pkg/sitemap"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	flagRemoveBaseUrl = false
)

var ConvertSitemapToUrllistCommand = &cobra.Command{
	Use:   "tools:convert-sitemap-to-urllist [sitemapPath]",
	Short: "Convert a sitemap to an url list and print it to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("You have to specify exactly one url or file path to a sitemap xml\n" +
				"\tUsage: crab tools:convert-sitemap-to-urllist https://domain.com/sitemap.xml")
			os.Exit(1)
		}

		sitemapPath := args[0]

		urls, err := sitemap.FetchUrlsFromPath(sitemapPath, &http.Client{Timeout: 30 * time.Second})
		if err != nil {
			fmt.Printf("Could not read sitemap from %s\n\t%s\n", sitemapPath, err)
			os.Exit(1)
		}

		if flagRemoveBaseUrl {
			urls = removeBaseUrls(urls)
		}

		for _, url := range urls {
			fmt.Println(url)
		}
	},
}

func removeBaseUrls(urls []string) []string {
	newUrls := make([]string, len(urls))

	for i, oldUrl := range urls {
		u, err := url.Parse(oldUrl)
		if err != nil {
			continue
		}

		query := u.RawQuery
		if query != "" {
			query = "?" + query
		}

		fragment := u.Fragment
		if fragment != "" {
			fragment = "#" + fragment
		}

		newUrls[i] = u.Path + query + fragment
	}

	return newUrls
}

func init() {
	ConvertSitemapToUrllistCommand.PersistentFlags().BoolVarP(
		&flagRemoveBaseUrl,
		"remove-base-url",
		"",
		false,
		"remove base url from urls",
	)
}
