package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stderr, "./cloclwall Tokyo=localhost:8000 London=localhost:8001")
		os.Exit(1)
	}

	clockWall(os.Args...)
}

func clockWall(args ...string) {
	done := make([]chan string, len(args[1:]))
	for i, arg := range args[1:] {
		pair := strings.Split(arg, "=")
		if len(pair) != 2 {
			fmt.Fprintf(os.Stderr, "./cloclwall Tokyo=localhost:8000 London=localhost:8001")
			os.Exit(1)
		}
		host := pair[1]
		done[i] = make(chan string)
		go dial(host, done[i])
	}
	for {
		for i, arg := range args[1:] {
			pair := strings.Split(arg, "=")
			name := pair[0]
			fmt.Printf("%s %s", name, <-done[i])
		}
	}
}

func dial(host string, done chan<- string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	defer conn.Close()

	read := bufio.NewReader(conn)
	for {
		str, err := read.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		done <- str
	}
}
