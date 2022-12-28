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
	Use:   "crawl:sitemap [sitemapPath]",
	Short: "Crawl through a sitemap",
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

		client := http.Client{Timeout: sitemapCommandFlags.HttpTimeout}

		sitemapModifiers := crawler.RequestModifier{}
		applySitemapModifiers(&sitemapModifiers, sitemapCommandFlags)

		urls, err := sitemap.FetchUrlsFromPath(sitemapPath, &client, &sitemapModifiers)
		if err != nil {
			fmt.Printf("Could not read sitemap from %s\n\t%s\n", sitemapPath, err)
			os.Exit(1)
		}

		err = crawlUrls(urls, modifier, sitemapCommandFlags, cmd.OutOrStdout())
		if err != nil {
			fmt.Printf("Could not create crawlable URLs:\n\t%s\n", err)
			os.Exit(1)
		}
	},
}

func applySitemapModifiers(modifier *crawler.RequestModifier, flagOptions crawlerFlagOptions) {
	modifier.With(addUserAgentToRequest())

	if len(flagOptions.AuthUsername) > 0 || len(flagOptions.AuthPassword) > 0 {
		modifier.With(addHttpBasicAuthToRequest(flagOptions))
	}

	if len(flagOptions.CookieStrings) > 0 {
		modifier.With(addCookiesToRequest(flagOptions))
	}

	if len(flagOptions.HeaderStrings) > 0 {
		modifier.With(addHeadersToRequest(flagOptions))
	}
}

func init() {
	registerStandardCrawlCommandFlags(SitemapCommand, &sitemapCommandFlags)
}
