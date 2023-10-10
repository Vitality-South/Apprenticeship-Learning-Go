package mathutil

import "testing"

func TestIsEven(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Input  int
		Expect bool
	}{
		{
			Input:  0,
			Expect: true,
		},
		{
			Input:  -1,
			Expect: false,
		},
		{
			Input:  -2,
			Expect: true,
		},
		{
			Input:  1,
			Expect: false,
		},
		{
			Input:  2,
			Expect: true,
		},
		{
			Input:  3,
			Expect: false,
		},
		{
			Input:  -3,
			Expect: false,
		},
	}

	for _, test := range tests {
		got := IsEven(test.Input)

		if got != test.Expect {
			t.Errorf("IsEven(%d) == %t, want %t", test.Input, got, test.Expect)
		}
	}
}

func TestIsOdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Input  int
		Expect bool
	}{
		{
			Input:  0,
			Expect: false,
		},
		{
			Input:  -1,
			Expect: true,
		},
		{
			Input:  -2,
			Expect: false,
		},
		{
			Input:  1,
			Expect: true,
		},
		{
			Input:  2,
			Expect: false,
		},
		{
			Input:  3,
			Expect: true,
		},
		{
			Input:  -3,
			Expect: true,
		},
	}

	for _, test := range tests {
		got := IsOdd(test.Input)

		if got != test.Expect {
			t.Errorf("IsOdd(%d) == %t, want %t", test.Input, got, test.Expect)
		}
	}
}
