package ascii

import "testing"

func TestNormalizeArtWidth(t *testing.T) {

	cases := []struct {
		in  []string
		out []string
	}{
		{[]string{"a", "abc"}, []string{"a  ", "abc"}},
		{[]string{"abc", "a"}, []string{"abc", "a  "}},
		{[]string{"", "a"}, []string{" ", "a"}},
		{[]string{"hi", "go", "a"}, []string{"hi", "go", "a "}},
		{[]string{"same", "size"}, []string{"same", "size"}},
		{[]string{"x"}, []string{"x"}},
		{[]string{}, []string{}},
		{[]string{"ab", "abcd"}, []string{"ab  ", "abcd"}},
		{[]string{"1", "22", "333"}, []string{"1  ", "22 ", "333"}},
		{[]string{"space "}, []string{"space "}},
	}

	for i, c := range cases {
		got := NormalizeArtWidth(c.in)
		for j := range got {
			if got[j] != c.out[j] {
				t.Errorf("case %d row %d expected %q got %q", i, j, c.out[j], got[j])
			}
		}
	}
}