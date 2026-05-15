package main

import "testing"

func TestPadArtRows(t *testing.T) {

	cases := []struct {
		rows  []string
		width int
		out   []string
	}{
		{[]string{"hi", "go"}, 4, []string{"hi  ", "go  "}},
		{[]string{"abc"}, 2, []string{"abc"}},
		{[]string{""}, 3, []string{"   "}},
		{[]string{"a"}, 1, []string{"a"}},
		{[]string{"a"}, 0, []string{"a"}},
		{[]string{"ab", "c"}, 3, []string{"ab ", "c  "}},
		{[]string{"longer"}, 3, []string{"longer"}},
		{[]string{"x", "yy"}, 5, []string{"x    ", "yy   "}},
		{[]string{}, 5, []string{}},
		{[]string{"go"}, -1, []string{"go"}},
	}

	for i, c := range cases {
		got := PadArtRows(c.rows, c.width)
		for j := range got {
			if got[j] != c.out[j] {
				t.Errorf("case %d row %d expected %q got %q", i, j, c.out[j], got[j])
			}
		}
	}
}