package main

import "testing"

func TestTrimArtRows(t *testing.T) {

	cases := []struct {
		in  []string
		out []string
	}{
		{[]string{"hi  ", "go  "}, []string{"hi", "go"}},
		{[]string{"a "}, []string{"a"}},
		{[]string{"   "}, []string{""}},
		{[]string{""}, []string{""}},
		{[]string{"  a  "}, []string{"  a"}},
		{[]string{"abc"}, []string{"abc"}},
		{[]string{"x   ", " y "}, []string{"x", " y"}},
		{[]string{"\t "}, []string{"\t"}}, // tab preserved, only space removed
		{[]string{"  go  ", "lang   "}, []string{"  go", "lang"}},
		{[]string{}, []string{}},
	}

	for i, c := range cases {
		got := TrimArtRows(c.in)
		for j := range got {
			if got[j] != c.out[j] {
				t.Errorf("case %d row %d expected %q got %q", i, j, c.out[j], got[j])
			}
		}
	}
}
