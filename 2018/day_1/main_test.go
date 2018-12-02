package main

import (
	"testing"
)

func TestReduce(t *testing.T) {
	tests := []testCase{
		testCase{
			in: in{
				numbers: []int{0, 0, 0},
				agg:     0,
			},
			out: 0,
		},
		testCase{
			in: in{
				numbers: []int{0, 1, 0},
				agg:     0,
			},
			out: 1,
		},
		testCase{
			in: in{
				numbers: []int{0, 0, -1},
				agg:     0,
			},
			out: -1,
		},
		testCase{
			in: in{
				numbers: []int{-5, 1, 1},
				agg:     0,
			},
			out: -3,
		},
	}

	for _, test := range tests {
		if res := reduce(test.in.numbers, test.in.agg); res != test.out {
			t.Errorf("Expected %d, got %d", test.out, res)
		}
	}
}

type testCase struct {
	in  in
	out int
}

type in struct {
	numbers []int
	agg     int
}
