package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	var c tempconv.Celsius = 1.0
	var f tempconv.Fahrenheit = 1.0
	var k tempconv.Kelvin = 1.0

	fmt.Println(c.String())
	fmt.Println(f.String())
	fmt.Println(k.String())

	fmt.Printf("%v = %v\n", c.String(), tempconv.CToF(c))
	fmt.Printf("%v = %v\n", c.String(), tempconv.CToK(c))
	fmt.Printf("%v = %v\n", f.String(), tempconv.FToC(f))
	fmt.Printf("%v = %v\n", f.String(), tempconv.FToK(f))
	fmt.Printf("%v = %v\n", k.String(), tempconv.KToF(k))
	fmt.Printf("%v = %v\n", k.String(), tempconv.KToC(k))
}
