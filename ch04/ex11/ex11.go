package main

import (
	"fmt"
	"os"

	"./github"
)

const usage = `./ex11 command <repository_name> <owner_name> [<issue_number>]
	command:
		create
		get
		edit
		close
    Environement variable: TOKEN (GitHub token)
`

func main() {
	if len(os.Args) < 4 || len(os.Args) > 6 {
		fmt.Println(usage)
		os.Exit(0)
	}
	if os.Getenv("TOKEN") == "" {
		fmt.Println(usage)
		os.Exit(0)
	}

	var issue string
	if len(os.Args) == 5 {
		issue = os.Args[4]
	}
	command := os.Args[1]
	repository := os.Args[2]
	owner := os.Args[3]

	token := os.Getenv("TOKEN")

	c := github.NewClient(token)

	var err error
	switch command {
	case "get":
		err = c.GetIssue(owner, repository, issue)
	case "create":
		err = c.CreateIssue(owner, repository)
	case "edit":
		err = c.EditIssue(owner, repository, issue)
	case "close":
		err = c.CloseIssue(owner, repository, issue)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
