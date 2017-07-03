package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
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

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	in := make(chan string)

	// Read input
	go func() {
		for input.Scan() {
			in <- input.Text()
		}
	}()

	tick := time.NewTicker(1 * time.Second)

	cnt := 0
	for {
		select {
		case str := <-in:
			go echo(c, str, 1*time.Second)
			cnt = 0
		case <-tick.C:
			fmt.Println(cnt)
			cnt++
		default:
		}
		if cnt == 10 {
			fmt.Println("connection closed")
			break
		}
	}
	tick.Stop()
	return
	// for input.Scan() {
	// 	echo(c, input.Text(), 1*time.Second)
	// }
}
