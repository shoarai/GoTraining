// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Issueshtml prints an HTML table of issues matching the search terms.
// Example:
// $ go buid github.com/shoarai/GoTraining/ch4/ex14/main.go
// $ ./main repo:golang/go commenter:gopherbot json encoder
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

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
	Milestone Milestone
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Number      int
	Title       string
	HTMLURL     string `json:"html_url"`
	Description string
}

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
	<th>Title</th>
  <th>Milestone</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	<td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var milestoneList = template.Must(template.New("milestoneList").Parse(`
<h1>milestones</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
	<th>Title</th>
  <th>Description</th>
</tr>
{{range .Items}}
<tr>
  <td>{{.Milestone.Number}}</td>
	<td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
	<td>{{.Milestone.Description}}</td>
</tr>
{{end}}
</table>
`))

var userList = template.Must(template.New("userList").Parse(`
<h1>users</h1>
<table>
<tr style='text-align: left'>
	<th>Name</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}

	if err := milestoneList.Execute(w, result); err != nil {
		log.Fatal(err)
	}

	if err := userList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}

// SearchIssues queries the GitHub issue tracker.
func searchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
