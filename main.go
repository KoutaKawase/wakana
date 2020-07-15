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
	var outputPath string
	nameOption := c.String("name")

	if directoryPath := c.String("output"); directoryPath != "" {
		outputPath = directoryPath + GetFileName(downloadURL, nameOption)
	} else {
		outputPath = fmt.Sprintf("./%s", GetFileName(downloadURL, nameOption))
	}

	err := downloadFile(outputPath, downloadURL)

	if err != nil {
		cli.Exit(err, 87)
	}

	fmt.Println("Downloaded from " + downloadURL)

	return nil
}

//GetFileName 絶対パスのファイルURLからファイル名を取得する
func GetFileName(downloadURL string, nameOption string) string {
	//ユーザー指定のファイルネームオプションがあればそのままそれを返す
	if nameOption != "" {
		return nameOption
	}

	//一番最後がスラッシュで終わっていればそのスラッシュを除いたものを取得
	if downloadURL[len(downloadURL)-1:] == "/" {
		downloadURL = downloadURL[:len(downloadURL)-1]
	}
	slashLastIndex := strings.LastIndex(downloadURL, "/")
	//+1しないと /demo.goみたいにスラッシュ自身が入るので+1している
	fileName := downloadURL[slashLastIndex+1:]
	return fileName
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
		Name:  "Wakana",
		Usage: "single file downloader for GitHub.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Specify the output destination. example) /path/to/outputDirectory `directoryPath`",
			},
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "Specify the file name to use for the downloaded file.",
			},
		},
		Action: action,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
