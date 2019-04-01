package strutils

import "testing"

func TestTruncateString(t *testing.T) {
	cases := []struct {
		input string
		limit int
		want  string
	}{
		{"Hello, World!", 20, "Hello, World!"},
		{"Hello, World!", 13, "Hello, World!"},
		{"Hello, World!", 10, "Hello, Wor"},
		{"Hello, World!", 0, ""},
	}

	for _, c := range cases {
		got := TruncateString(c.input, c.limit)

		if got != c.want {
			t.Errorf("Truncate(%q, %d) failed. Wanted: '%q' but got: '%q'",
				c.input, c.limit, c.want, got)
		}
	}
}
