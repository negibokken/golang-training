package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const usage = `./ex01 <port>
  port: port number (default 8000)

  Environment variable: TZ
    ex)  TZ=Tokyo/Asia(default)
         TZ=US/Eastern
         TZ=Europe/London
`

func main() {
	timeZone := "Tokyo/Asia"
	port := "8000"
	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stderr, "%v", usage)
		os.Exit(1)
	}
	if os.Args[1] != "" {
		port = os.Args[1]
	}
	if tz := os.Getenv("TZ"); tz != "" {
		timeZone = tz
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatal(err)
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, loc)
	}
}

func handleConn(c net.Conn, location *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
