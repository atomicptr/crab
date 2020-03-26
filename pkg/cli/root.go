package cli

import (
	"github.com/atomicptr/crab/pkg/cli/crawl"
	"github.com/atomicptr/crab/pkg/meta"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     "crab",
	Short:   "Crab - A versatile web crawler for all your needs.",
	Version: meta.VersionString(),
}

func init() {
	rootCommand.AddCommand(crawl.Command)
	rootCommand.AddCommand(crawl.SitemapCommand)
}
