package main

import (
	"fmt"
	"time"
)

func main() {

	Start(2, 3)
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	for l := range m {
		fmt.Printf("\n\nstuck in for loop %d \n\n", l)
		s := Karthic{}
		s.test = l
		CollectTasks(s)
	}
	StopAllWorkers(2)
	time.Sleep(3 * time.Second)

}
