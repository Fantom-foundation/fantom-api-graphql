package gvclear

import "fmt"

// main implements the entry point into the Governance Clearing
// tools used to resolve pending Governance tasks.
func main() {
	// collect configuration options
	cfg := config()

	// connect the Opera/RPC to be able to interact with it
	con, err := connect(cfg)
	if err != nil {
		fmt.Printf("Can not connect Opera node!\n%s", err.Error())
		return
	}

	// get the governance contract interface
	gov, err := governance(cfg, con)
	if err != nil {
		fmt.Printf("Can not access Governance contract!\n%s", err.Error())
		return
	}
}
