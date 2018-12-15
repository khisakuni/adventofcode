package main

import (
	"testing"
)

func TestFindRepeat(t *testing.T) {
	type testCase struct {
		in  []int
		out int
	}
	tests := []testCase{
		testCase{
			in:  []int{1, -1},
			out: 0,
		},
		testCase{
			in:  []int{1, -2, 3, 1, 1, -2},
			out: 2,
		},
		testCase{
			in:  []int{3, 3, 4, -2, -4},
			out: 10,
		},
		testCase{
			in:  []int{-6, 3, 8, 5, -6},
			out: 5,
		},
		testCase{
			in:  []int{7, 7, -2, -7, -4},
			out: 14,
		},
	}
	for _, test := range tests {
		if res := repeat(test.in); res != test.out {
			t.Errorf("Expected %d, got %d", test.out, res)
		}
	}
}

func TestReduce(t *testing.T) {
	type in struct {
		numbers []int
		agg     int
	}
	type testCase struct {
		in  in
		out int
	}
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
