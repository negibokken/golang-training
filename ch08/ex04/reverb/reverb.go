package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	done := make(chan struct{})
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		// Echo
		go func(c net.Conn, shout string, delay time.Duration) {
			defer wg.Done()
			fmt.Fprintln(c, "\t", strings.ToUpper(shout))
			time.Sleep(delay)
			fmt.Fprintln(c, "\t", shout)
			time.Sleep(delay)
			fmt.Fprintln(c, "\t", strings.ToLower(shout))
			time.Sleep(delay)
			done <- struct{}{}
		}(c, input.Text(), 1*time.Second)
	}

	go func() {
		wg.Wait()
		close(done)
	}()
}
