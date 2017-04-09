package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	messages := []string{}
	for range os.Args[1:] {
		message := <-ch
		fmt.Println(message)
		messages = append(messages, message)
	}
	saveToFile(messages)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %vb", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func saveToFile(messages []string) error {
	date := time.Now().String()
	date = strings.Replace(date, " ", "_", -1)
	fileName := fmt.Sprintf("%v.txt", date)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("File open error: %v", err)
	}
	for _, b := range messages {
		_, err := file.WriteString(fmt.Sprintf("%v\n", b))
		if err != nil {
			return err
		}
	}
	return nil
}
