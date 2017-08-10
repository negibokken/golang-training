package main

import (
	"fmt"
	"os"
	"time"
)

type pingpong struct {
	in  chan struct{}
	out chan struct{}
}

func ping(p pingpong, cancel chan struct{}) {
	fmt.Println("ping start")
pingloop:
	for {
		p.out <- struct{}{}
		// fmt.Println("ping")
		<-p.in
		pingCnt++
		if cancelled() {
			break pingloop
		}
	}
}

func pong(p pingpong, cancel chan struct{}) {
	fmt.Println("pong start")
pongloop:
	for {
		<-p.in
		p.out <- struct{}{}
		pongCnt++
		// fmt.Println("pong")
		if cancelled() {
			break pongloop
		}
	}
}

func cancelled() bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}

var pingCnt = 0
var pongCnt = 0
var cancel = make(chan struct{})

func main() {
	p := make(chan struct{})
	q := make(chan struct{})

	go ping(pingpong{p, q}, cancel)
	go pong(pingpong{q, p}, cancel)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()

	timeout := time.Second * 1
	timer := time.AfterFunc(timeout, func() {
		fmt.Println("1 second, finished")
		close(cancel)
	})

loop:
	for {
		if cancelled() {
			fmt.Println("finish")
			break loop
		}
	}
	timer.Stop()
	fmt.Printf("ping sent %v times struct{}.\n", pingCnt)
	fmt.Printf("pong sent %v times struct{}.\n", pongCnt)
}
