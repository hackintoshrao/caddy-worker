package main

import (
	"fmt"
)

// A buffered channel that we can send work requests on.
var WorkQueue chan Task

func CollectTasks(task Task) {
	// Push the work onto the queue.
	WorkQueue <- task
	fmt.Println("Work request queued")
	// And let the user know their work request was created.
	return
}

func Start(nWorkers int, taskBufferSize int) {
	WorkQueue = make(chan Task, taskBufferSize)

	fmt.Println("Starting the dispatcher")
	StartDispatcher(nWorkers)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")

}
