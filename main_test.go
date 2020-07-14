package main

import "testing"

var downloadURL string = "https://github.com/mattn/go-gtk/blob/master/_example/demo/demo.go"

func TestGitHubファイルURLがblobからrawに書き変えられているか(t *testing.T) {
	converted := ConvertURL(downloadURL)
	expected := "https://github.com/mattn/go-gtk/raw/master/_example/demo/demo.go"

	if converted != expected {
		t.Errorf("欲しいのは %s。 でも %s だった", expected, converted)
	}
}

func Test有効なGitHubURLならtrueを返すか(t *testing.T) {
	isValid := IsValidURL(downloadURL)

	if isValid != true {
		t.Errorf("IsValidURL(%s) は trueが返ってこなければなりません", downloadURL)
	}
}

func Test有効なワードが含まれてないURLならfalseを返すか(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"hogehoge", false},
		{"https://example.com", false},
		{"htt", false},
		{"https://github.com/KoutaKawase", false},
		{"https://blob/", false},
		{"うおおおおおおｄさｋｆｊｋｓｄｌｊｆ", false},
	}

	for _, c := range cases {
		got := IsValidURL(c.in)

		if got != c.want {
			t.Errorf("IsValidURL(%s) == %t, want %t", c.in, got, c.want)
		}
	}
}

func Test最後のファイル名のみを切り出せているか(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{downloadURL, "demo.go"},
		{"https://github.com/mattn/go-gtk/blob/master/_example/demo/demo", "demo"},
		{"https://github.com/mattn/go-gtk/blob/master/_example/demo/fugafuga.py/", "fugafuga.py"},
		{"https://github.com/mattn/hoge/blob/master/fuga/piyo.test.ts", "piyo.test.ts"},
	}

	for _, c := range cases {
		got := GetFileName(c.in)

		if got != c.want {
			t.Errorf("GetFileName(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}
