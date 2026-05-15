package main

import "testing"

func TestCountVisibleChars(t *testing.T) {

	cases := []struct {
		in  []string
		out int
	}{
		{[]string{"A A", "###"}, 5},
		{[]string{"   "}, 0},
		{[]string{"abc"}, 3},
		{[]string{"a b c"}, 3},
		{[]string{}, 0},
		{[]string{"123"}, 3},
		{[]string{"! @ #"}, 3},
		{[]string{"go lang"}, 6},
		{[]string{"x", "y", "z"}, 3},
		{[]string{"a a a", "b b"}, 5},
	}

	for i, c := range cases {
		got := CountVisibleChars(c.in)
		if got != c.out {
			t.Errorf("case %d expected %d got %d", i, c.out, got)
		}
	}
}
