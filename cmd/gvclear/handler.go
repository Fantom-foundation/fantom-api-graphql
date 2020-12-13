package main

import (
	"fmt"
	"math/big"
)

// handleTasks processes the pending tasks on Governance contract.
func handleTasks(con *connector) error {
	// get the number of tasks
	tasks, err := con.gov.TasksCount(defaultCallOpts(&con.key.From, con.client))
	if err != nil {
		fmt.Printf("Can not count pending tasks!\n%s\n", err.Error())
		return err
	}

	// log what we have
	count := tasks.Uint64()
	fmt.Printf("Found %d pending tasks.\n", count)

	// anything to do?
	if 0 == count {
		fmt.Println("Nothing to do ...")
		return nil
	}

	// process them in batches
	for i := uint64(0); i < count; i += con.batchLen {
		doHandleTasks(i, con)
	}

	// wait for the processing to finish
	return nil
}

// doHandleTasks handles a batch of tasks starting with the specified index.
func doHandleTasks(from uint64, con *connector) {
	// execute the transaction
	tx, err := con.gov.HandleTasks(con.key, new(big.Int).SetUint64(from), new(big.Int).SetUint64(con.batchLen))
	if err != nil {
		fmt.Printf("Can not handle tasks for index %d!\n%s\n", from, err.Error())
		return
	}

	// inform about the trx
	fmt.Printf("Handle transaction: %s\n", tx.Hash().String())
}

// cleanupTasks processes the task cleanup on Governance contract.
func cleanupTasks(con *connector) error {
	// execute the transaction
	tx, err := con.gov.TasksCleanup(con.key, new(big.Int).SetUint64(con.batchLen))
	if err != nil {
		fmt.Printf("Can not cleanup tasks!\n%s\n", err.Error())
		return err
	}

	// inform about the trx
	fmt.Printf("Cleanup transaction: %s\n", tx.Hash().String())
	return nil
}
