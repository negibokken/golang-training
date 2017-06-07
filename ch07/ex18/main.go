package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func makeTree(dec *xml.Decoder) (Node, error) {
	var stack []*Element
	var result *Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			elem := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, elem)
			}
			stack = append(stack, elem)
			if result == nil {
				result = elem
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}
	return result, nil
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	result, err := makeTree(dec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	fmt.Printf("%#v", result)

}
