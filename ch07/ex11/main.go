package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var db database

func main() {
	db = database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	fmt.Println("Listening localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for k, v := range db {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		fmt.Fprintf(w, "invalid parameter. Exaple: /create?item=xxx&price=xxx\n")
		return
	}
	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "%s already exists.\n", item)
		return
	}
	pricef, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "price is invalid.\n")
		return
	}
	db[item] = dollars(pricef)
	fmt.Fprintf(w, "%s: %s is saved.\n", item, price)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		fmt.Fprintf(w, "invalid parameter. Example: /read?item=xxx\n")
		return
	}
	pricef, ok := db[item]
	if !ok {
		fmt.Fprintf(w, "%s does not exist", item)
		return
	}
	fmt.Fprintf(w, "%s's price is %s\n", item, pricef)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		fmt.Fprintf(w, "invalid parameter. Example: /update?item=xxx&price=xxx\n")
		return
	}
	_, ok := db[item]
	if !ok {
		fmt.Fprintf(w, "%s does not exist", item)
		return
	}
	pricef, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	db[item] = dollars(pricef)
	fmt.Fprintf(w, "%s's price is %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		fmt.Fprintf(w, "invalid parameter. Example: /delete?item=xxx")
		return
	}
	_, ok := db[item]
	if !ok {
		fmt.Fprintf(w, "%s does not exist", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "%s is deleted", item)
}
