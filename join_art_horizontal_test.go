package ascii

import "testing"

func TestJoinArtHorizontal(t *testing.T) {

	cases := []struct {
		left  []string
		right []string
		out   []string
	}{
		{[]string{"A"}, []string{"1"}, []string{"A1"}},
		{[]string{"A", "B"}, []string{"1", "2"}, []string{"A1", "B2"}},
		{[]string{"hi"}, []string{"go"}, []string{"higo"}},
		{[]string{"x", "y", "z"}, []string{"1", "2", "3"}, []string{"x1", "y2", "z3"}},
		{[]string{}, []string{}, []string{}},
		{[]string{"a"}, []string{"b"}, []string{"ab"}},
		{[]string{"go"}, []string{"lang"}, []string{"golang"}},
		{[]string{"A ", "B "}, []string{"1", "2"}, []string{"A 1", "B 2"}},
		{[]string{"x", "yy"}, []string{"1", "22"}, []string{"x1", "yy22"}},
		{[]string{"hi", "there"}, []string{"!!", "??"}, []string{"hi!!", "there??"}},
	}

	for i, c := range cases {
		got := JoinArtHorizontal(c.left, c.right)
		for j := range got {
			if got[j] != c.out[j] {
				t.Errorf("case %d row %d expected %q got %q", i, j, c.out[j], got[j])
			}
		}
	}
}