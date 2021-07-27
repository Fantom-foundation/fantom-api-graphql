// Package build keeps information about the app building
// environment. It uses liker ldflags to inject correct values
// into the building process. Please use Makefile to get correct
// results, or check it to inject the variables manually.
package build

import (
	"fantom-api-graphql/internal/config"
	"fmt"
	"runtime"
)

// Version represents the version of the app.
var Version = "undefined"

// Commit represents the GitHub commit hash the app was built from.
var Commit = "undefined"

// CommitTime represents the GitHub commit time stamp the app was built from.
var CommitTime = "undefined"

// Time represents the time of the app build.
var Time = "undefined"

// Compiler represents the information about the compiler used to build the app.
var Compiler = "undefined"

// Reset represents a token for terminal color reset.
var Reset = "\033[0m"

// Blue represents a token for blue terminal color setup.
var Blue = "\033[34m"

// init initializes the build reference on the given OS
func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Blue = ""
	}
}

// PrintVersion prints the version information
// into the std output.
func PrintVersion(cfg *config.Config) {
	fmt.Printf("%sApp Name:%s\t%s\n", Blue, Reset, cfg.AppName)
	fmt.Printf("%sApp Version:%s\t%s\n", Blue, Reset, Version)
	fmt.Printf("%sCommit Hash:%s\t%s\n", Blue, Reset, Commit)
	fmt.Printf("%sCommit Time:%s\t%s\n", Blue, Reset, CommitTime)
	fmt.Printf("%sBuild Time:%s\t%s\n", Blue, Reset, Time)
	fmt.Printf("%sBuild By:%s\t%s\n", Blue, Reset, Compiler)
}

// Short returns a short, single line version of the app.
func Short(cfg *config.Config) string {
	return fmt.Sprintf("%s v%s, commit:%s, build:%s", cfg.AppName, Version, Commit, Time)
}
