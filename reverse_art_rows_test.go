package ascii

import "testing"

func TestReverseArtRows(t *testing.T) {

	cases := []struct {
		in  []string
		out []string
	}{
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"1"}, []string{"1"}},
		{[]string{}, []string{}},
		{[]string{"top", "mid", "bot"}, []string{"bot", "mid", "top"}},
		{[]string{"x", "y"}, []string{"y", "x"}},
		{[]string{"go", "lang"}, []string{"lang", "go"}},
		{[]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}},
		{[]string{"same"}, []string{"same"}},
		{[]string{"1", "2"}, []string{"2", "1"}},
		{[]string{"hi", "there"}, []string{"there", "hi"}},
	}

	for i, c := range cases {
		got := ReverseArtRows(c.in)
		for j := range got {
			if got[j] != c.out[j] {
				t.Errorf("case %d row %d expected %q got %q", i, j, c.out[j], got[j])
			}
		}
	}
}