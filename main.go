package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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
	outputPath := "./demo.go"

	err := downloadFile(outputPath, downloadURL)

	if err != nil {
		cli.Exit(err, 87)
	}

	fmt.Println("Downloaded: " + downloadURL)

	return nil
}

func downloadFile(filepath string, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	out, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
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
