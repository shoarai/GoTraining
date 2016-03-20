// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/shoarai/GoTraining/ch4/ex12/xkcd"
)

func print(comic *xkcd.Comic) {
	data, err := json.MarshalIndent(&comic, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func main() {
	var len int
	fmt.Printf("Load the number of comics: ")
	fmt.Scan(&len)
	xkcd.LoadComics(len)

	for {
		var num int
		fmt.Printf("Comic number: ")
		fmt.Scan(&num)
		comic := xkcd.GetComic(num)
		if comic != nil {
			print(comic)
		} else {
			fmt.Println("Not loaded comic")
		}
	}
}