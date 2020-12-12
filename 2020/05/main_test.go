package main

import "testing"

func TestParse(t *testing.T) {
	tests := []struct{
		in string
		row float64
		col float64
		id float64
	}{
		{
			in:"FBFBBFFRLR",
			row: 44,
			col: 5,
			id: 357,
		},
		{
			in:"BFFFBBFRRR",
			row: 70,
			col: 7,
			id: 567,
		},
		{
			in:"FFFBBBFRRR",
			row: 14,
			col: 7,
			id: 119,
		},
		{
			in:"BBFFBBFRLL",
			row: 102,
			col: 4,
			id: 820,
		},
	}
	for _, c := range tests {
		row, col, id := parse(c.in)
		if row != c.row {
			t.Errorf("%s: expected row %f, got %f", c.in, c.row, row)
		}
		if col != c.col {
			t.Errorf("%s: expected col %f, got %f", c.in, c.col, col)
		}
		if id != c.id {
			t.Errorf("%s: expected id %f, got %f", c.in, c.id, id)
		}
	}
}
