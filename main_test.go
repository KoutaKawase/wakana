package main

import "testing"

func TestGitHubファイルURLがblobからrawに書き変えられているか(t *testing.T) {
	url := ConvertURL("https://github.com/mattn/go-gtk/blob/master/_example/demo/demo.go")
	expected := "https://github.com/mattn/go-gtk/raw/master/_example/demo/demo.go"

	if url != expected {
		t.Errorf("欲しいのは %s。 でも %s だった", expected, url)
	}
}
