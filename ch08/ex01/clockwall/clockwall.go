package main

import (
	"fmt"
	"io"
	"log"
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
	for _, arg := range args[1:] {
		pair := strings.Split(arg, "=")
		if len(pair) == 1 {
			fmt.Fprintf(os.Stderr, "./cloclwall Tokyo=localhost:8000 London=localhost:8001")
			os.Exit(1)
		}
		conn, err := net.Dial("tcp", pair[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		defer conn.Close()
		go handleConn(conn, pair[0])
	}
}

func handleConn(c net.Conn, location string) {
	for {
		fmt.Fprintf(os.Stdout, "Location: %v", location)
		mustCopy(os.Stdout, c)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
