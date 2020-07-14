package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func action(c *cli.Context) error {
	if c.NArg() == 0 {
		fmt.Println("No URL argument!")
		return nil
	}
	url := c.Args().Get(0)
	fmt.Println(url)
	return nil
}

//ConvertURL 与えられたURLをダウンロード可能なものに変換する
func ConvertURL(url string) string {
	return strings.Replace(url, "/blob/", "/raw/", 1)
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
