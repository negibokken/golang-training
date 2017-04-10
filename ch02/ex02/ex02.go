package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"./length"
)

func main() {

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	var m length.Meter = length.Meter(num)
	var i length.Inch = length.Inch(num)
	var f length.Feet = length.Feet(num)

	fmt.Println(m.String())
	fmt.Println(i.String())
	fmt.Println(f.String())

	fmt.Printf("%v = %v\n", m.String(), length.MToI(m))
	fmt.Printf("%v = %v\n", m.String(), length.MToF(m))
	fmt.Printf("%v = %v\n", i.String(), length.IToM(i))
	fmt.Printf("%v = %v\n", i.String(), length.IToF(i))
	fmt.Printf("%v = %v\n", f.String(), length.FToM(f))
	fmt.Printf("%v = %v\n", f.String(), length.FToI(f))
}
