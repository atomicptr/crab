package meta

import "fmt"

var Version = ""
var GitCommit = ""
var UserAgentBase = "github.com/atomicptr/crab"

// UserAgent returns the user agent
func UserAgent() string {
	return fmt.Sprintf("%s (%s)", UserAgentBase, VersionString())
}

// VersionString returns the build version and commit.
func VersionString() string {
	commitString := ""
	// ignore warning, this value will be later added as a build flag
	if len(GitCommit) >= 7 {
		commitString = fmt.Sprintf("-%s", GitCommit[:7])
	}

	if Version == "" {
		Version = "dev"
	}

	return fmt.Sprintf("%s%s", Version, commitString)
}
