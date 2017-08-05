package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
)

// ListResponse is response of `go list -json`
type ListResponse struct {
	Dir         string
	ImportPath  string
	Name        string
	Target      string
	Stale       bool
	StateReason string
	Root        string
	GoFiles     []string
	Imports     []string
	Deps        []string
}

var out io.Writer = os.Stdout

func execGoList(arg string) ([]string, error) {
	args := []string{"list", "-json", `-f={{.ImportPath}} {{join .Deps " "}}`, arg}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		return nil, err
	}
	data := new(ListResponse)
	json.Unmarshal(out, data)
	return data.Deps, nil
}

func printPackages(ma map[string]bool) {
	var keys []string
	for k := range ma {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(out, "\t- %s\n", k)
	}
}

var all bool

func init() {
	flag.BoolVar(&all, "a", false, "summarize all package")
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "./ex04 <package_or_dir>")
		os.Exit(1)
	}
	arg := flag.Args()[0]
	fmt.Fprintf(out, "%s\n", arg)
	pkgs, err := execGoList(arg)
	if err != nil {
		log.Fatal(err)
	}
	var packages = make(map[string]bool)
	sort.Strings(pkgs)
	for _, pkg := range pkgs {
		// Print deps package
		packages[pkg] = true
		if all {
			fmt.Fprintf(out, "\t- %s:\n", pkg)
		}

		// Print each deps package of deps package
		ps, err := execGoList(pkg)
		if err != nil {
			log.Fatal(err)
		}
		sort.Strings(ps)
		for _, p := range ps {
			packages[p] = true
			if all {
				fmt.Fprintf(out, "\t\t- %s\n", p)
			}
		}
	}
	if !all {
		printPackages(packages)
	}
}
