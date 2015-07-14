package main

import "net/http"

type WorkRequest struct {
	Name   string
	Writer http.ResponseWriter
	Done   chan bool
}
