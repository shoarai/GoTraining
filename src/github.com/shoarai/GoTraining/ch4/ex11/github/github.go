// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"strconv"
	"time"
)

const (
	APIURL    = "https://api.github.com"
	IssuesURL = APIURL + "/search/issues"
)

type Auth struct {
	UserName string
	Password string
}

type Repository struct {
	Owner string
	Repo  string
}

type IssueCreateRequest struct {
	Title     string
	Body      string
	Assignee  string
	Milestone int
	Labels    []string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func IssueCreateURL(r *Repository) string {
	return APIURL + "/repos/" + r.Owner + "/" + r.Repo + "/issues"
}

func IssueGetURL(r *Repository, num int) string {
	return APIURL + "/repos/" + r.Owner + "/" + r.Repo + "/issues/" + strconv.Itoa(num)
}
