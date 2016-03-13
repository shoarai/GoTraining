// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateIssue creates the GitHub issue.
func CreateIssue(repo *Repository, issue *IssueCreateRequest, auth *Auth) {
	url := IssueCreateURL(repo)
	json, _ := json.Marshal(issue)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return
		// return nil, err
	}
	req.Header.Add(authHeader(auth))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	// resp, err := http.Post(CreateIssuesURL, bodyType string, body io.Reader)
}

// GetIssue gets the GitHub issue.
func GetIssue(repo *Repository, number int, auth *Auth) (*Issue, error) {
	url := IssueGetURL(repo, number)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(authHeader(auth))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func authHeader(auth *Auth) (string, string) {
	s := auth.UserName + ":" + auth.Password
	return "Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(s))
}
