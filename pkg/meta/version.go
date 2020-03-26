package meta

import "fmt"

var Version = "dev"
var GitCommit = ""
var UserAgent = "github.com/atomicptr/crab"

func VersionString() string {
	commitString := ""
	// ignore warning, this value will be later added as a build flag
	if len(GitCommit) >= 7 {
		commitString = fmt.Sprintf(" (%s)", GitCommit[:7])
	}

	return fmt.Sprintf("%s%s", Version, commitString)
}
