package main

import (
	"log"
	"os"

	"github.com/negibokken/golang-training/ch13/ex04/bzip"
)

func main() {
	w, err := bzip.NewWriter(os.Stdout)
	if err != nil {
		log.Println("hello1")
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Println("hello")
		log.Fatalf("bzipper: close: %v\n", err)
	}
}
