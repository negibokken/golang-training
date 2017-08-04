package main

import (
	"fmt"
	arc "github.com/negibokken/golang-training/ch10/ex02"
	_ "github.com/negibokken/golang-training/ch10/archive/zip"
	_ "github.com/negibokken/golang-training/ch10/archive/tar"
)

func printArchive(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := arc.Open(f)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: arc FILE ...")
	}
	exitCode := 0
	for _, filename := range os.Args[1:] {
		err := printArchive(filename)
		if err != nil {
			log.Print(err)
			exitCode=2
		}
	}
	os.Exit(exitCode)
}
