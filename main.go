package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func action(c *cli.Context) error {
	if c.NArg() == 0 {
		fmt.Println("No URL argument!")
		return nil
	}

	downloadURL := c.Args().Get(0)

	if !IsValidURL(downloadURL) {
		return cli.Exit("Invalid argument. Please input URL.", 86)
	}

	downloadURL = ConvertURL(downloadURL)
	return nil
}

//IsValidURL github.com/.../blob/...じゃないURLを弾く
func IsValidURL(downloadURL string) bool {
	//ParseではなくParseRequestURIなのは絶対パスしか想定してないから
	u, err := url.ParseRequestURI(downloadURL)

	if err != nil {
		return false
	}

	if u.Host == "github.com" && strings.Contains(u.Path, "/blob/") {
		return true
	}

	return false
}

//ConvertURL 与えられたURLをダウンロード可能なものに変換する
func ConvertURL(downloadURL string) string {
	return strings.Replace(downloadURL, "/blob/", "/raw/", 1)
}

func main() {
	app := &cli.App{
		Name:   "Wakana",
		Usage:  "single file downloader for GitHub.",
		Action: action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
