// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

/*
	Using a select statement, add a timeout to the echo server from Section 8.3 so
	that is disconnects any client that shouts nothing within 10 seconds
*/

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	defer c.Close()
	abort := time.After(10 * time.Second)
	input := bufio.NewScanner(c)
	text := make(chan string)
	go func() {
		for {
			if input.Scan() {
				text <- input.Text()
			}

		}

	}()
outer:
	for {
		select {
		case <-abort:
			break outer
		case x := <-text:
			go echo(c, x, 1*time.Second)
			abort = time.After(10 * time.Second)
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
