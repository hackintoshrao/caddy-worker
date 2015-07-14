package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"time"
)

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)

func Collector(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	work := &WorkRequest{}
	// Parse the delay.
	err := json.NewDecoder(r.Body).Decode(work)
	if err != nil {
		http.Error(w, "Bad Name value: "+err.Error(), http.StatusBadRequest)
		return
	}
	work.Writer = w
	// Now, we take the delay, and the person's name, and make a WorkRequest out of them.

	// Push the work onto the queue.
	WorkQueue <- *work
	fmt.Println("Work request queued")

	<-work.Done
	fmt.Println("unblock")
	// And let the user know their work request was created.
	//w.WriteHeader(http.StatusCreated)
	return
}
