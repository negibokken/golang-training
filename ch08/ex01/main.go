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
    ex)  TZ=Tokyo/Asia
         TZ=US/Eastern(default)
         TZ=Europe/London
`

func main() {
	timeZone := "US/Eastern"
	port := "8000"
	if len(os.Args) >= 2 {
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
	fmt.Printf("%s: localhost:%s listening", timeZone, port)
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
