package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

/*
	Construct a pipeline that connects an arbitrary number of goroutines with channels.
	What is the maximum number of pipeline stages that you can create without running
	out of memory? How long does a value take to transit the entire pipeline?
*/

func main() {
	if len(os.Args) != 2 {
		log.Fatal("missing number of goroutines")
	}
	num, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		log.Fatalf("problem parsing number: %v", err)
	}

	result := make(chan struct{})
	t := time.Now()
	go myFunc(num, result)
	<-result
	fmt.Printf("%d goroutines took %s\n", num, time.Since(t))
}

func myFunc(count int64, result chan struct{}) {
	if count == 0 {
		result <- struct{}{}
		return
	}
	ch := make(chan struct{})
	go myFunc(count-1, ch)
	result <- <-ch
}
