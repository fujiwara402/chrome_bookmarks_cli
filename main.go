package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fujiwara402/chrome_bookmarks_cli/parser"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	BookmarksFile = "/home/chrml/.config/google-chrome/Default/Bookmarks"

	root = kingpin.
		Flag("root", "Select root directory. Choose from 'BookmarkBar', 'Synced' or 'Other'.").
		Default("BookmarkBar").
		Short('r').
		Enum("BookmarkBar", "Synced", "Other")

	dir = kingpin.
		Flag("dir", "Select sub directory.").
		Short('d').
		String()

	tree = kingpin.
		Flag("tree", "show as tree").
		Short('t').
		Bool()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	err := _main()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func _main() error {
	raw, err := ioutil.ReadFile(BookmarksFile)
	if err != nil {
		return err
	}

	bs, err := parser.Perse(raw)
	if err != nil {
		return err
	}

	switch *root {
	case "BookmarkBar":
		bs.Root.BookmarkBar.PrintItems()
	case "Synced":
		bs.Root.Synced.PrintItems()
	case "Other":
		bs.Root.Other.PrintItems()
	default:
		return errors.New("予期せぬエラーが発生しました")
	}

	return nil
}
