package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/negibokken/golang-training/ch12/ex12/params"
)

func checkPass(v interface{}) error {
	num, ok := v.(int)
	if !ok {
		return fmt.Errorf("not an int: %v", v)
	}
	if num <= 30 {
		return fmt.Errorf("%d <= 30", num)
	}
	return nil
}

func main() {

	type unpacked struct {
		Post int `http:p,check:"check"`
	}

	var out unpacked
	req := &http.Request{Form: url.Values{"p": []string{"100"}}}
	// req := &http.Request{Form: url.Values{"pass": []string{"wrong password"}}}
	checks := map[string]params.Check{
		"checkPass": checkPass,
	}

	err := params.Unpack(req, &out, checks)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("params:%v\n", out)
}
