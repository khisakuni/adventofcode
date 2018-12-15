package main

import "testing"

func TestHasRep(t *testing.T) {
	type in struct {
		num int
		str string
	}
	type testCase struct {
		in  in
		out bool
	}

	tests := []testCase{
		testCase{
			in: in{
				str: "",
				num: 2,
			},
			out: false,
		},
		testCase{
			in: in{
				str: "a",
				num: 2,
			},
			out: false,
		},
		testCase{
			in: in{
				str: "ab",
				num: 2,
			},
			out: false,
		},
		testCase{
			in: in{
				str: "aa",
				num: 2,
			},
			out: true,
		},
		testCase{
			in: in{
				str: "aba",
				num: 2,
			},
			out: true,
		},
		testCase{
			in: in{
				str: "aaba",
				num: 2,
			},
			out: false,
		},
	}

	for _, test := range tests {
		if res := hasRep(test.in.str, test.in.num); res != test.out {
			t.Errorf("Expected %v for %+v, got %v", test.out, test.in, res)
		}
	}
}
