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
		fmt.Printf("Cleanup failed!")
		return
	}

	// process tasks
	if err := processTasks(con); err != nil {
		fmt.Printf("Cleanup failed!")
		return
	}

	// cleanup tasks

	// close connection

	// we are done here
	fmt.Printf("Done.")
}

// processTasks processes the pending tasks on Governance contract.
func processTasks(con *connector) error {
	// get the number of tasks
	tasks, err := con.gov.TasksCount(defaultCallOpts(&con.key.From, con.client))
	if err != nil {
		fmt.Printf("Can not count pending tasks!\n%s", err.Error())
		return err
	}

	// log what we have
	fmt.Printf("Found %d pending tasks.", tasks.Uint64())
	return nil
}
