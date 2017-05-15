package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	day := map[int]github.Issue{}
	week := map[int]github.Issue{}
	month := map[int]github.Issue{}
	year := map[int]github.Issue{}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now().UTC()

	// classify
	for _, item := range result.Items {
		diffHour := now.Sub(item.CreatedAt).Hours()
		if diffHour < 24 {
			day[item.Number] = *item
		} else if diffHour < 24*7 {
			week[item.Number] = *item
		} else if diffHour < 24*30 {
			month[item.Number] = *item
		} else {
			year[item.Number] = *item
		}
	}
	printItem(day, "day")
	printItem(week, "week")
	printItem(month, "month")
	printItem(year, "more than year ago")
}

func printItem(m map[int]github.Issue, title string) {
	fmt.Printf("--- %v ----\n", title)
	for _, item := range m {
		fmt.Printf("#%-5d %9.9s %.55s %s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

}
