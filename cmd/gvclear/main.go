package main

import (
	"fmt"
)

// main implements the entry point into the Governance Clearing
// tools used to resolve pending Governance tasks.
func main() {
	// collect configuration options
	cfg := config()

	// attach the node and prepare to process tasks
	con, err := attach(cfg)
	if err != nil {
		fmt.Println("Cleanup failed!")
		return
	}

	// process tasks
	if err := handleTasks(con); err != nil {
		fmt.Println("Cleanup failed!")
		return
	}

	// cleanup tasks
	if err := cleanupTasks(con); err != nil {
		fmt.Println("Cleanup failed!")
		return
	}

	// we are done here
	fmt.Println("Done.")
}
