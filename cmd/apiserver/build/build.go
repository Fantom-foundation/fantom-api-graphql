// Package build keeps information about the app building
// environment. It uses liker ldflags to inject correct values
// into the building process. Please use Makefile to get correct
// results, or check it to inject the variables manually.
package build

import "fmt"

var Version = "undefined"
var Commit = "undefined"
var Time = "undefined"
var Compiler = "undefined"

// PrintVersion prints the version information
// into the std output.
func PrintVersion() {
	fmt.Printf("Fantom GraphQL API Server\n")
	fmt.Printf("App Version:\t%s\n", Version)
	fmt.Printf("Build Commit:\t%s\n", Commit)
	fmt.Printf("Build Time:\t%s\n", Time)
	fmt.Printf("Build GoLang:\t%s\n", Compiler)
}

// Short returns a short, single line version of the app.
func Short() string {
	return fmt.Sprintf("v%s, commit %s (%s by %s)", Version, Commit, Time, Compiler)
}
