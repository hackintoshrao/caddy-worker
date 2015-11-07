package main

import "fmt"

//stop finite workers , increase workers , increase buffer, real time numbers on worker status
var WorkerQueue chan chan Task
var QuitChanQueue chan chan bool
var EndDispatcher chan bool

func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan Task, nworkers)
	QuitChanQueue = make(chan chan bool, nworkers)
	EndDispatcher = make(chan bool, 1)
	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue, QuitChanQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work request")
				// go func() {
				worker := <-WorkerQueue

				fmt.Println("Dispatching work request")
				worker <- work
				//}()
			case <-EndDispatcher:
				fmt.Println("Closing Dispatcher")
				return
			}
		}
	}()
}

func StopAllWorkers(nWorkers int) {
	var quitChan chan bool
	EndDispatcher <- true

	for i := 0; i < nWorkers; i++ {
		quitChan = <-QuitChanQueue
		go func(quitChan chan bool) {
			quitChan <- true
		}(quitChan)
	}

}
