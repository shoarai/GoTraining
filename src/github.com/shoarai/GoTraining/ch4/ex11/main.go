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
		"shoarai", "Dammy",
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
	case "edit":
		edit()
	default:
		fmt.Println("Input command")
	}
}

func create() {
	var issue github.IssueCreateRequest
	issue.Title = "TestTitle1"
	issue.Body = "TestBody"
	github.CreateIssue(&repo, &issue, &auth)
}

func get() {
	var num int
	fmt.Println("Issue number: ")
	fmt.Scan(&num)

	issue, err := github.GetIssue(&repo, num, &auth)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("#%-5d %9.9s %.55s %s\n",
		issue.Number, issue.User.Login, issue.Title, issue.Body)
}

func edit() {
	var num int
	fmt.Println("Issue number: ")
	fmt.Scan(&num)

	currentIssue, err := github.GetIssue(&repo, num, &auth)
	if err != nil {
		log.Fatal(err)
		return
	}

	var issue github.IssueEditRequest
	issue.Title = currentIssue.Title
	issue.Body = currentIssue.Body
	issue.State = "closed"
	err = github.EditIssue(&repo, num, &issue, &auth)
	if err != nil {
		fmt.Println(err)
	}
}
