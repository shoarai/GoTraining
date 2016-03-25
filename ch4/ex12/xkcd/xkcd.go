// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Comic is type of comic information
type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
}

var comics map[int]*Comic

// GetComic gets the comic infomation
func GetComic(num int) *Comic {
	comic, ok := comics[num]
	if !ok {
		return nil
	}
	return comic
}

// LoadComics gets the Comics
func LoadComics(num int) {
	comics = make(map[int]*Comic)
	for i := 1; i < num+1; i++ {
		comic, err := loadComic(i)
		if err != nil {
			continue
		}
		comics[i] = comic
		fmt.Printf("loaded: %d / %d\n", i, num)
	}
}

func loadComic(num int) (*Comic, error) {
	url := getComicURL(num)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}

	return &comic, nil
}

func getComicURL(num int) string {
	return "https://xkcd.com/" + strconv.Itoa(num) + "/info.0.json"
}
