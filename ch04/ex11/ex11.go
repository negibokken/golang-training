package main

import (
	"flag"
	"fmt"
	"os"

	"./github"
)

const usage = `
./ex11 command -repo <repository_name> -owner <owner_name> [-issue <number>]
	command:
		create
		get
		edit
		close
`

func main() {
	r := flag.String("repo", "", "repository name")
	o := flag.String("owner", "", "owner name")
	i := flag.String("issue", "", "issue number")

	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()

	// command := flag.Args()[0]
	// for _, cm := range flag.Args() {
	// 	fmt.Println(cm)
	// }
	// fmt.Println(command)
	// if command == "" {
	// 	flag.Usage()
	// 	os.Exit(0)
	// }

	fmt.Println(*r)
	fmt.Println(*o)
	fmt.Println(*i)
	command := "aaa"

	fmt.Println(flag.NFlag())
	if *i != "" && flag.NFlag() != 2 {
		flag.Usage()
		os.Exit(0)
	}

	token := os.Getenv("TOKEN")

	c := github.NewClient(token)
	r = r
	c = c
	o = o

	switch command {
	case "get":
		// c.GetIssue()
	case "create":
	case "edit":

	case "close":
	}

}
