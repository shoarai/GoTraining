// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Omdapi outputs movie poster as jpg.
// Example:
// $ go buid github.com/shoarai/GoTraining/ch4/ex13/main.go
// $ ./main Frozen > poster.jpg
package main

import (
	"fmt"
	"os"

	"github.com/shoarai/GoTraining/ch4/ex13/omdbapi"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input movie title as argument")
		return
	}
	title := os.Args[1]

	movie, err := omdbapi.GetMovie(title)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Getting movie failed: %s\n", err)
		os.Exit(1)
	}

	if movie.Title == "" {
		fmt.Fprintf(os.Stderr, "Movie title of %q is not existed\n", title)
		os.Exit(1)
	}

	err = omdbapi.GetPoster(os.Stdout, movie)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Getting poster failed: %s\n", err)
		os.Exit(1)
	}
}
