package crawl

import (
	"fmt"
	"os"

	"github.com/atomicptr/crab/pkg/crawler"
	"github.com/spf13/cobra"
)

var crawlCommandFlags = newCrawlerFlagOptions()

var Command = &cobra.Command{
	Use:   "crawl [domains...]",
	Short: "Crawl singular URLs",
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateCrawlerFlagOptions(crawlCommandFlags); err != nil {
			fmt.Printf("Flag options are invalid:\n\t%s\n", err)
			os.Exit(1)
		}

		if len(args) == 0 {
			fmt.Println("You have to specify at least one URL.\n\tUsage: crab crawl https://github.com")
			os.Exit(1)
		}

		modifier := crawler.RequestModifier{}
		registerStandardCrawlCommandFlagModifiers(&modifier, crawlCommandFlags)

		err := crawlUrls(args, modifier, crawlCommandFlags)
		if err != nil {
			fmt.Printf("Could not create crawlable URLs:\n\t%s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	registerStandardCrawlCommandFlags(Command, &crawlCommandFlags)
}
