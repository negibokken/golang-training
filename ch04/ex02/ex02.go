package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	h := flag.String("hash", "SHA256", "SHA256, SHA384 or SHA512")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("./ex01 -hash <SHA_Type> <str1>")
		os.Exit(0)
	}
	arg := flag.Args()[0]
	switch *h {
	case "SHA256":
		c1 := sha256.Sum256([]byte(arg))
		fmt.Printf("%v:\n  %v =>\n\t%x\n", *h, arg, c1)
	case "SHA384":
		c1 := sha512.Sum384([]byte(arg))
		fmt.Printf("%v:\n  %v =>\n\t%x\n", *h, arg, c1)
	case "SHA512":
		c1 := sha512.Sum512([]byte(arg))
		fmt.Printf("%v:\n  %v =>\n\t%x\n", *h, arg, c1)
	}
}
