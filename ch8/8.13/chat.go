// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

/*
	Make the chat server disconnect idle lients, such as those that have sent no messages
	in the last five minutes. Hint: calling conn.Close() in another goroutine unblocks active
	Read calls such as the one done by input.Scan()
*/

type client struct {
	ch   chan<- string
	name string
} // an outgoing message channel

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
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{
		ch:   ch,
		name: who,
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli
	disconnectTimer := time.NewTimer(1 * time.Minute)
	go func() {
		select {
		case <-disconnectTimer.C:
			conn.Close()
		}
	}()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		disconnectTimer.Reset(1 * time.Minute)
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
