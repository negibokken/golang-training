package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

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

type client struct {
	user string
	msg  chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.msg <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			for cl := range clients {
				cli.msg <- cl.user + "is online"
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{who, ch}
	input := bufio.NewScanner(conn)

	timeout := time.Minute * 5
	timer := time.AfterFunc(timeout, func() {
		fmt.Fprintln(conn, "timeout, connection closed")
		conn.Close()
	})
	for input.Scan() {
		timer.Stop()
		messages <- who + ": " + input.Text()
		timer = time.AfterFunc(timeout, func() {
			fmt.Fprintln(conn, "timeout, connection closed")
			conn.Close()
		})
	}
	timer.Stop()

	leaving <- client{who, ch}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for cli := range ch {
		fmt.Fprintln(conn, cli)
	}
}
