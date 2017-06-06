package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type attr struct {
	local string
	value string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "./ex17 <id or class> <value>")
		os.Exit(1)
	}
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	var attrStack []attr
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
			for _, t := range tok.Attr {
				if t.Name.Local == os.Args[1] && t.Value == os.Args[2] {
					fmt.Printf("\nattr: %s, value: %s\n", t.Name.Local, t.Value)
				}
				attrStack = append(attrStack, attr{t.Name.Local, t.Value})
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if hasTarget(attrStack, os.Args[1], os.Args[2]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				attrStack = attrStack[:len(attrStack)-1]
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func hasTarget(attrStack []attr, id, value string) bool {
	for _, a := range attrStack {
		if a.local == id && a.value == value {
			return true
		}
	}
	return false
}
