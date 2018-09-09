package parser

import (
	"encoding/json"
	"fmt"
)

type Bookmark struct {
	Root     Root   `json:"roots"`
	Checksum string `json:"checksum"`
	Version  int    `json:"version"`
}

type Root struct {
	BookmarkBar            Item   `json:"bookmark_bar"`
	Synced                 Item   `json:"synced"`
	Other                  Item   `json:"other"`
	SyncTransactionVersion string `json:"sync_transaction_version"`
}

type Item struct {
	Items                  []Item   `json:"children"`
	ID                     string   `json:"id"`
	Name                   string   `json:"name"`
	URL                    string   `json:"url"`
	Type                   string   `json:"type"`
	DateAdded              string   `json:"date_added"`
	SyncTransactionVersion string   `json:"sync_transaction_version"`
	MetaInfo               MetaInfo `json:"meta_info"`
}

type MetaInfo struct {
	LastVisitedDesktop string `json:"last_visited_desktop"`
}

func (c Item) PrintItems() {
	switch c.Type {
	case "url":
		fmt.Println("⭐", c.Name, ":", c.URL)
	case "folder":
		fmt.Println()
		fmt.Println("📂", c.Name)
		for _, ch := range c.Items {
			ch.PrintItems()
		}
	}
}

func Perse(raw []byte) (*Bookmark, error) {
	bs := new(Bookmark)
	if err := json.Unmarshal(raw, bs); err != nil {
		return nil, err
	}

	return bs, nil
}
