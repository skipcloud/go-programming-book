package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Write a program with two goroutines that send messages back and forth
	over two unbuffered channels in ping-pong fashion. How many communications
	per second can the program sustain
*/

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	timer := time.NewTimer(5 * time.Second)

	var wg sync.WaitGroup
	var count int64

	// ping
	wg.Add(1)
	go func() {
	loop:
		for val := range ch1 {
			select {
			case ch2 <- val:
				count += 1
			case <-timer.C:
				break loop
			}
		}
		// close the channels outside of the for loop,
		// this way if the timer finishes or the receive
		// channel is closed we close the other channels
		// receive channel, aka our sending channel
		close(ch2)
		wg.Done()
	}()

	// pong
	wg.Add(1)
	go func() {
	loop:
		for val := range ch2 {
			select {
			case ch1 <- val:
				count += 1
			case <-timer.C:
				break loop
			}
		}
		close(ch1)
		wg.Done()
	}()

	// kick off the ping pong match
	ch1 <- struct{}{}
	wg.Wait()
	fmt.Printf("over 5 seconds there was %d communications, or %d communications per second\n", count, count/5.0)
}
