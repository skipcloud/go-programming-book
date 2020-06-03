// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/*
	Failure of any client program to read data in a timely manner ultimately causes
	all clients to get stuck. Modify the broadcaster to skip a message rather than
	wait if a client writer is not ready to accept it. Alternatively, add buffering
	to each client's outgoing message channel so that most messages are not dropped;
	the broadcaster should use a non-blocking send to this channel.
*/

type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			if len(clients) > 0 {
				cli.ch <- "Also online:"
				for c := range clients {
					cli.ch <- c.name
				}
			} else {
				cli.ch <- "No one else is online\n"
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string, 10) // buffered outgoing client messages
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)
	fmt.Fprintf(conn, "Enter your name: ")
	input.Scan()
	who := input.Text()
	cli := client{
		ch:   ch,
		name: who,
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
