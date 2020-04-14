package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sort"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

/*
	Modify clock2 to accept a port number, and write a program, clockwall,
	that acts as a client of several clock servers at once, reading the times
	from each one and displaying the results in a table, akin to the wall
	of clocks seen in some business offices. If you have access to geographically
	distributed computers, run instances remotely; otherwise run local
	instances on different ports with fake timezones.

	$ TZ=US/Eastern ./clock2 -port 8010 &
	$ TZ=Asia/Tokyo ./clock2 -port 8020 &
	$ TZ=Europe/London ./clock2 -port 8030 &
	$ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
*/

func main() {
	args := parseInput()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "missing arguments")
	}

	// For consistent formatting when printing to Stdout
	// create a list of the city names and sort them
	cities := make([]string, 0, len(args))
	for k, _ := range args {
		cities = append(cities, k)
	}
	sort.Strings(cities)

	// times holds onto the output from the TCP connection
	// for each city, a RWMutex helps stop any panic from
	// concurrent writes to the map
	times := map[string]string{}
	var mutex sync.RWMutex
	for city, addr := range args {
		city := city
		addr := addr
		go func() {
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				log.Print(err)
				return
			}
			defer conn.Close()
			for {
				t, err := bufio.NewReader(conn).ReadString('\n')
				if err == io.EOF {
					return
				} else if err != nil {
					log.Print(err)
					return
				}
				mutex.Lock()
				times[city] = strings.TrimRight(t, "\n")
				mutex.Unlock()
				time.Sleep(1 * time.Second)
			}
		}()
	}

	for {
		writeToStdOut(cities, times)
		time.Sleep(1 * time.Second)
	}
}

func parseInput() map[string]string {
	args := map[string]string{}
	for _, val := range os.Args[1:] {
		// split London=localhost:8000
		s := strings.Split(val, "=")
		args[s[0]] = s[1]
	}
	return args
}

func writeToStdOut(cities []string, times map[string]string) {
	const format = "%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 5, 2, ' ', 0)

	fmt.Fprintf(tw, format, "City", "Local Time")
	fmt.Fprintf(tw, format, "----", "----------")

	for _, city := range cities {
		fmt.Fprintf(tw, format, city, times[city])
	}
	tw.Flush()
}
