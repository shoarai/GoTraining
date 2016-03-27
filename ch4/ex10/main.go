// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var monthIss github.IssuesSearchResult
	var yearIss github.IssuesSearchResult
	var beforeYearIss github.IssuesSearchResult
	now := time.Now()
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)
	for _, item := range result.Items {
		if item.CreatedAt.After(month) {
			monthIss.Items = append(monthIss.Items, item)
			monthIss.TotalCount++
		} else if item.CreatedAt.After(year) {
			yearIss.Items = append(yearIss.Items, item)
			yearIss.TotalCount++
		} else {
			beforeYearIss.Items = append(beforeYearIss.Items, item)
			beforeYearIss.TotalCount++
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("%d issues within a month:\n", monthIss.TotalCount)
	for _, item := range monthIss.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d issues over a month and within a year:\n", yearIss.TotalCount)
	for _, item := range yearIss.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d issues over a year:\n", beforeYearIss.TotalCount)
	for _, item := range beforeYearIss.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

/*
//!+textoutput
$ go build github.com/shoarai/GoTraining/ch4/ex10/main.go
$ ./main repo:golang/go is:open json decoder
14 issues:
2 issues within a month:
#14750 cyberphon encoding/json: parser ignores the case of member names
#14640 AntiPaste encoding/json: decoding a null value does not replace p
3 issues over a month and within a year:
#11046     kurin encoding/json: Decoder internally buffers full input
#12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
#13558  ajwerner io: MultiReader should be more efficient when chained m
9 issues over a year:
#5680    eaigner encoding/json: set key converter on en/decoder
#8658  gopherbot encoding/json: use bufio
#5901        rsc encoding/json: allow override type marshaling
#6716  gopherbot encoding/json: include field name in unmarshal error me
#7872  extempora encoding/json: Encoder internally buffers full output
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
