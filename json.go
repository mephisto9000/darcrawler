package main

import (
	"encoding/json"
	"io/ioutil"
)

type PageInfo struct {
	Title    string   `json:"title"`
	Keywords []string `json:"keywords"`
}

func NewPageInfo(title string, keywords []string) PageInfo {
	return PageInfo{title, keywords}
}

func exportToFile(infos []PageInfo, filename string) error {
	data, err := json.Marshal(infos)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// func importFromFile(filename string) ([]PageInfo, error) {
// 	bytes, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var infos []PageInfo
// 	if err := json.Unmarshal(bytes, infos); err != nil {
// 		return nil, err
// 	}
// 	return infos, nil
// }
