package crawl

import (
	"fmt"
	"github.com/atomicptr/crab/pkg/crawler"
	"github.com/atomicptr/crab/pkg/urllist"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var listCommandFlags = newCrawlerFlagOptions()

var ListCommand = &cobra.Command{
	Use:   "crawl:list [url-list.txt]",
	Short: "Crawl through a list of URLs",
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateCrawlerFlagOptions(listCommandFlags); err != nil {
			fmt.Printf("Flag options are invalid:\n\t%s\n", err)
			os.Exit(1)
		}

		if len(args) != 1 {
			fmt.Println("You have to specify exactly one url or file path to an url list\n" +
				"\tUsage: crab crawl:list ./urls.txt")
			os.Exit(1)
		}

		modifier := crawler.RequestModifier{}
		registerStandardCrawlCommandFlagModifiers(&modifier, listCommandFlags)

		listPath := args[0]

		urls, err := urllist.FetchUrlsFromPath(listPath, &http.Client{Timeout: listCommandFlags.HttpTimeout})
		if err != nil {
			fmt.Printf("Could not read url list from %s\n\t%s\n", listPath, err)
			os.Exit(1)
		}

		err = crawlUrls(urls, modifier, listCommandFlags, cmd.OutOrStdout())
		if err != nil {
			fmt.Printf("Could not create crawlable URLs:\n\t%s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	registerStandardCrawlCommandFlags(ListCommand, &listCommandFlags)
}
