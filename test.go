package main

import (
	"fmt"
	//"time"
)

type Karthic struct {
	test int
}

func (k Karthic) ExecuteTask() interface{} {
	fmt.Printf("\nInside task without sleep %d\n", k.test)
	//time.Sleep(3 * time.Second)
	return k.test
}
