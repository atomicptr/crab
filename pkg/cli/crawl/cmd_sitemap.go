package crawl

import (
	"fmt"
	"github.com/atomicptr/crab/pkg/crawler"
	"github.com/atomicptr/crab/pkg/sitemap"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var sitemapCommandFlags = newCrawlerFlagOptions()

var SitemapCommand = &cobra.Command{
	Use: "crawl:sitemap [sitemapPath]",
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateCrawlerFlagOptions(sitemapCommandFlags); err != nil {
			fmt.Printf("Flag options are invalid:\n\t%s\n", err)
			os.Exit(1)
		}

		if len(args) != 1 {
			fmt.Println("You have to specify exactly one url or file path to a sitemap xml\n" +
				"\tUsage: crab crawl:sitemap https://domain.com/sitemap.xml")
			os.Exit(1)
		}

		modifier := crawler.RequestModifier{}
		registerStandardCrawlCommandFlagModifiers(&modifier, sitemapCommandFlags)

		sitemapPath := args[0]

		urls, err := sitemap.FetchUrlsFromPath(sitemapPath, &http.Client{Timeout: sitemapCommandFlags.HttpTimeout})
		if err != nil {
			fmt.Printf("Could not read sitemap from %s\n\t%s\n", sitemapPath, err)
			os.Exit(1)
		}

		err = crawlUrls(urls, modifier, sitemapCommandFlags)
		if err != nil {
			fmt.Printf("Could not create crawlable URLs:\n\t%s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	registerStandardCrawlCommandFlags(SitemapCommand, &sitemapCommandFlags)
}
