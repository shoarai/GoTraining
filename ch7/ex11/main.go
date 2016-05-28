// Copyright © 2016 shoarai

// Http is an e-commerce server that registers
// the /list, /price, /create, /update and /delete
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: $%d\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "$%d\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	newPriceStr := req.URL.Query().Get("price")
	newPrice, ok := strconv.Atoi(newPriceStr)
	if ok != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "price error\n")
		return
	}

	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "existing item: %q\n", item)
	} else {
		if regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9]+`).Match([]byte(item)) {
			db[item] = newPrice
			fmt.Fprintf(w, "created item: %q\n", item)
		} else {
			fmt.Fprintf(w, "item name error: %q\n", item)
		}
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	newPriceStr := req.URL.Query().Get("price")
	newPrice, ok := strconv.Atoi(newPriceStr)
	if ok != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "price error\n")
		return
	}

	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		db[item] = newPrice
		fmt.Fprintf(w, "$%d\n", newPrice)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "delted item: %q\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
