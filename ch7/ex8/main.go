// Copyright © 2016 shoarai

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	custom := customSort{tracks, []string{"Title", "Artist", "Album", "Year", "Length"}}

	for {
		var key string
		fmt.Printf("\nInput sort key %v :\n", custom.keys)
		fmt.Scan(&key)

		err := custom.selectKey(key)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Custom:")
		sort.Sort(custom)

		printTracks(tracks)
	}
}

type customSort struct {
	t    []*Track
	keys []string
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.lessByKeys(x.t[i], x.t[j], x.keys) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x *customSort) selectKey(key string) error {
	for i, v := range x.keys {
		if v == key {
			x.keys = append(x.keys[:i], x.keys[i+1:]...)
			x.keys = append([]string{v}, x.keys...)
			return nil
		}
	}
	return fmt.Errorf("key error")
}

func (customSort) lessByKeys(x, y *Track, keys []string) bool {
	for _, v := range keys {
		switch v {
		case "Title":
			if x.Title != y.Title {
				return x.Title < y.Title
			}
		case "Artist":
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
		case "Album":
			if x.Album != y.Album {
				return x.Album < y.Album
			}
		case "Year":
			if x.Year != y.Year {
				return x.Year < y.Year
			}
		case "Length":
			if x.Length != y.Length {
				return x.Length < y.Length
			}
		}
	}
	return false
}
