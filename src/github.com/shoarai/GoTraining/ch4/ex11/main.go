// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shoarai/GoTraining/ch4/ex11/github"
)

var repo github.Repository
var auth github.Auth

func main() {
	if len(os.Args[1:]) < 0 {
		return
	}

	repo = github.Repository{
		Owner: "shoarai",
		Repo:  "Dammy",
	}

	fmt.Println("User name: ")
	fmt.Scan(&auth.UserName)
	fmt.Println("Password: ")
	fmt.Scan(&auth.Password)

	switch os.Args[1] {
	case "create":
		create()
	case "get":
		get()
	default:
		fmt.Println("Input command")
	}
}

func create() {
	var issue github.IssueCreateRequest
	issue.Title = "TestTitle"
	issue.Body = "TestBody"
	github.CreateIssue(&repo, &issue, &auth)
}

func get() {
	var num int
	fmt.Println("Issue number: ")
	fmt.Scan(&num)

	item, err := github.GetIssue(&repo, num, &auth)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("#%-5d %9.9s %.55s %s\n",
		item.Number, item.User.Login, item.Title, item.Body)
}
