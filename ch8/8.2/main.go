package main

import (
	"fmt"
	"log"
	"net"

	"github.com/skipcloud/go-programming-book/ch8/8.2/ftpserver"
)

/*
	Implement a concurrent File Transfer Protocol (FTP) server. The server should
	interpret commands from each client such as cd to change directory, ls to list
	a directory, get to send the contents of a file, and close to close the
	connection. You can use the standard ftp command as the client, or write your own


	note from Skip: here is the FTP RFC, it's very handy for this exercise
					http://www.faqs.org/rfcs/rfc959.html

					The code isn't perfect but I got farther than I thought I would
*/

const addr = "localhost:8000"

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server running on %s\n", addr)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("got connection")
		go ftpserver.HandleConn(conn)
	}
}
