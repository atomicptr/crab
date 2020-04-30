package cli

import (
	"github.com/atomicptr/crab/pkg/cli/crawl"
	"github.com/atomicptr/crab/pkg/cli/tools"
	"github.com/atomicptr/crab/pkg/meta"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     "crab",
	Short:   "Crab - A versatile tool to crawl dozens of URLs from a given source.",
	Version: meta.VersionString(),
}

func init() {
	rootCommand.AddCommand(crawl.Command)
	rootCommand.AddCommand(crawl.SitemapCommand)
	rootCommand.AddCommand(crawl.ListCommand)

	rootCommand.AddCommand(tools.ConvertSitemapToUrllistCommand)
}
