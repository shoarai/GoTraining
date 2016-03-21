// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateIssue creates the GitHub issue.
func CreateIssue(repo *Repository, issue *IssueCreateRequest, auth *Auth) error {
	url := IssueCreateURL(repo)
	json, _ := json.Marshal(issue)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	req.Header.Add(authHeader(auth))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("create failed: %s", resp.Status)
	}
	return nil
}

// GetIssue gets the GitHub issue.
func GetIssue(repo *Repository, number int, auth *Auth) (*Issue, error) {
	url := IssueGetURL(repo, number)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(authHeader(auth))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// EditIssue edits the GitHub issue.
func EditIssue(repo *Repository, num int, issue *IssueEditRequest, auth *Auth) error {
	url := IssueEditURL(repo, num)
	json, _ := json.Marshal(issue)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	req.Header.Add(authHeader(auth))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("edit failed: %s", resp.Status)
	}
	return nil
}

func authHeader(auth *Auth) (string, string) {
	s := auth.UserName + ":" + auth.Password
	return "Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(s))
}
