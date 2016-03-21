// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/shoarai/GoTraining/ch4/ex11/github"
)

var repo github.Repository
var auth github.Auth

func main() {
	if len(os.Args) < 2 {
		printCommands()
		return
	}

	repo = github.Repository{
		"shoarai", "Dummy",
	}

	fmt.Printf("User name: ")
	fmt.Scan(&auth.UserName)
	fmt.Printf("Password: ")
	fmt.Scan(&auth.Password)

	switch os.Args[1] {
	case "create":
		create()
	case "get":
		get()
	case "edit":
		edit()
	default:
		printCommands()
	}
}

func printCommands() {
	fmt.Println(`The commands are:
	create: create new issue
	get:    get the issue
	edit:   edit the issue`)
}

func create() {
	var issue github.IssueCreateRequest
	in := bufio.NewReader(os.Stdin)

	fmt.Printf("Title: ")
	line, _, err := in.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	issue.Title = string(line)

	fmt.Printf("Body: ")
	line, _, err = in.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	issue.Body = string(line)

	err = github.CreateIssue(&repo, &issue, &auth)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("The issue has been created")
}

func get() {
	var num int
	fmt.Printf("Issue number: ")
	fmt.Scan(&num)

	issue, err := github.GetIssue(&repo, num, &auth)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Title: %s\n", issue.Title)
	fmt.Printf("State: %s\n", issue.State)
	fmt.Printf("User: %s\n", issue.User.Login)
	fmt.Printf("Body: %s\n", issue.Body)
}

func edit() {
	var num int
	fmt.Printf("Issue number: ")
	fmt.Scan(&num)

	currentIssue, err := github.GetIssue(&repo, num, &auth)
	if err != nil {
		log.Fatal(err)
		return
	}

	var issue github.IssueEditRequest
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("Current Title: %s\n", currentIssue.Title)
	fmt.Printf("    New Title: ")
	line, _, err := in.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	issue.Title = string(line)
	if issue.Title == "" {
		issue.Title = currentIssue.Title
	}

	fmt.Printf("Current Body: %s\n", currentIssue.Body)
	fmt.Printf("    New Body: ")
	lines, err := readLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	issue.Body = lines

	// line, _, err = in.ReadLine()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// issue.Body = string(line)
	if issue.Body == "" {
		issue.Body = currentIssue.Body
	}

	fmt.Printf("Current State: %s\n", currentIssue.State)
	fmt.Printf("    New State: ")
	line, _, err = in.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	issue.State = string(line)
	if issue.State == "" {
		issue.State = currentIssue.State
	}

	err = github.EditIssue(&repo, num, &issue, &auth)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The issue has been edited")
}

func readLines() (string, error) {
	var lines string
	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		lines += string(line) + "\n"
	}
	return lines, nil
}
