package crawl

import (
	"net/http"
	"net/url"

	"github.com/atomicptr/crab/pkg/crawler"
	"github.com/spf13/cobra"
)

func registerStandardCrawlCommandFlags(cmd *cobra.Command, flagOptions *crawlerFlagOptions) {
	cmd.PersistentFlags().IntVarP(
		&flagOptions.NumberOfWorkers,
		"num-workers",
		"",
		defaultNumberOfWorkers,
		"set number of workers for crawling",
	)
	cmd.PersistentFlags().StringVarP(
		&flagOptions.PrefixUrl,
		"prefix-url",
		"",
		"",
		"prefix/replace all request urls with this one",
	)
	cmd.PersistentFlags().DurationVarP(
		&flagOptions.HttpTimeout,
		"http-timeout",
		"",
		defaultHttpTimeout,
		"set http timeout for requests",
	)
	cmd.PersistentFlags().StringSliceVarP(
		&flagOptions.CookieStrings,
		"cookie",
		"",
		nil,
		"add cookies (as key=value pairs) to each request",
	)
	cmd.PersistentFlags().StringSliceVarP(
		&flagOptions.HeaderStrings,
		"header",
		"",
		nil,
		"add headers (as key=value pairs) to each request",
	)
}

func registerStandardCrawlCommandFlagModifiers(modifier *crawler.RequestModifier, flagOptions crawlerFlagOptions) {
	modifier.With(addUserAgentToRequest())

	if isValidUrl(flagOptions.PrefixUrl) {
		modifier.With(addPrefixUrlToRequest(flagOptions.PrefixUrl))
	}

	if len(flagOptions.CookieStrings) > 0 {
		modifier.With(addCookiesToRequest(flagOptions))
	}

	if len(flagOptions.HeaderStrings) > 0 {
		modifier.With(addHeadersToRequest(flagOptions))
	}
}

func crawlUrls(urls []string, modifier crawler.RequestModifier, flagOptions crawlerFlagOptions) error {
	requests, err := crawler.CreateRequestsFromUrls(urls, modifier)
	if err != nil {
		return err
	}

	crawl := crawler.Crawler{
		NumberOfWorkers: flagOptions.NumberOfWorkers,
		HttpClient: http.Client{
			Timeout: flagOptions.HttpTimeout,
		},
	}
	crawl.Crawl(requests)

	return nil
}

func isValidUrl(rawUrl string) bool {
	if rawUrl == "" {
		return false
	}

	_, err := url.Parse(rawUrl)
	return err == nil
}
