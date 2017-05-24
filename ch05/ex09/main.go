package main

import "os"
import "fmt"
import "regexp"
import "io/ioutil"

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}

	expanded := expand(string(buf),
		func(s string) string {
			return s + s
		})
	fmt.Printf("%v\n", expanded)
}

func expand(s string, f func(string) string) string {
	re := regexp.MustCompile("\\$(\\w+)")
	return re.ReplaceAllString(s, f("${1}"))
}
