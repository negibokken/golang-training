package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	comics := map[int]Comic{
		1: Comic{"test Title", "test"},
		2: Comic{"one piece", "gomu"},
		3: Comic{"naruto", "ninja"},
		4: Comic{"yugio", "duel"},
	}

	type Want struct {
		num   int
		comic Comic
	}
	tests := []struct {
		title string
		want  Want
	}{
		{"test Title", Want{1, Comic{"test Title", "test"}}},
		{"one piece", Want{2, Comic{"one piece", "gomu"}}},
		{"naruto", Want{3, Comic{"naruto", "ninja"}}},
		{"yugio", Want{4, Comic{"yugio", "duel"}}},
	}
	for _, test := range tests {
		num, comic := search(comics, test.title)
		if num != test.want.num {
			t.Errorf("got: %v, expected: %v", num, test.want.num)
		}
		if *comic != test.want.comic {
			t.Errorf("got: %#v, expected: %#v", comic, test.want.comic)
		}
	}
}

func TestComicURL(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{1, "http://xkcd.com/1/info.0.json"},
		{2, "http://xkcd.com/2/info.0.json"},
		{3, "http://xkcd.com/3/info.0.json"},
	}

	for _, test := range tests {
		if got := comicURL(test.num); got != test.want {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
}
