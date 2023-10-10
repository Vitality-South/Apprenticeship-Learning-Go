package stringutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Input  string
		Expect string
	}{
		{
			Input:  "",
			Expect: "",
		},
		{
			Input:  "Hello, world",
			Expect: "dlrow ,olleH",
		},
		{
			Input:  "Hello, 世界",
			Expect: "界世 ,olleH",
		},
	}

	for _, test := range tests {
		got := Reverse(test.Input)

		if got != test.Expect {
			t.Errorf("Reverse(%q) == %q, want %q", test.Input, got, test.Expect)
		}
	}
}
