// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package omdbapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Movie struct {
	Title  string
	Poster string
}

func getPosterURL(title string) string {
	return "http://www.omdbapi.com/?t=" + title + "&y=&plot=short&r=json"
}

func GetMovie(title string) (*Movie, error) {
	url := getPosterURL(title)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed %s", resp.Status)
	}

	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func GetPoster(w io.Writer, movie *Movie) error {
	url := movie.Poster
	if url == "" {
		return fmt.Errorf("poster url is empty")
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get poster url failed %s", resp.Status)
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
