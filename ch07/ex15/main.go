package main

import (
	"os"
	"strconv"

	"fmt"

	"bufio"

	"./eval"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Expr: ")
	for scanner.Scan() {
		exprString := scanner.Text()
		if exprString == "" {
			break
		}

		expr, err := eval.Parse(string(exprString))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid expression %s", exprString)
			os.Exit(1)
		}

		vars := map[eval.Var]bool{}
		err = expr.Check(vars)
		if err != nil {
			fmt.Printf("Invalid expression: %s\n", err)
			continue
		}

		env := eval.Env{}
		for v := range vars {
			fmt.Printf("%s = ", v)
			for scanner.Scan() {
				a := scanner.Text()
				arg, err := strconv.ParseFloat(string(a), 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Invalid value %s", err)
					os.Exit(1)
				}
				env[v] = arg
				break
			}
		}
		fmt.Printf("%f\n", expr.Eval(env))
		fmt.Print("Expr: ")
	}
}
