package main

import (
	"fmt"
)

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan Task, quitChanQueue chan chan bool) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:            id,
		Work:          make(chan Task),
		WorkerQueue:   workerQueue,
		QuitChan:      make(chan bool),
		QuitChanQueue: quitChanQueue,
	}

	return worker
}

type Worker struct {
	ID            int
	Work          chan Task
	WorkerQueue   chan chan Task
	QuitChan      chan bool
	QuitChanQueue chan chan bool
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w Worker) Start() {
	go func() {
		w.QuitChanQueue <- w.QuitChan
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work
			fmt.Printf("worker %d is Ready\n", w.ID)
			select {
			case task := <-w.Work:
				// Receive a work request.
				fmt.Printf("worker%d: Received work request\n", w.ID)
				task.ExecuteTask()

			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("\n\nworker %d stopping\n\n", w.ID)
				return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
// func (w Worker) Stop() {
// 	go func() {
// 		w.QuitChan <- true
// 	}()
// }
