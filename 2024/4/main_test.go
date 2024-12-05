package main

import "testing"

func TestSearch(t *testing.T) {
	for _, tc := range []struct {
		input [][]byte
		start coord
		want  int
	}{
		{
			input: [][]byte{
				{'X', 'M', 'A', 'S'},
			},
			start: coord{0, 0},
			want:  1,
		},
		{
			input: [][]byte{
				{'X'},
				{'M'},
				{'A'},
				{'S'},
			},
			start: coord{0, 0},
			want:  1,
		},
		{
			input: [][]byte{
				{'X', '.', '.', '.'},
				{'.', 'M', '.', '.'},
				{'.', '.', 'A', '.'},
				{'.', '.', '.', 'S'},
			},
			start: coord{0, 0},
			want:  1,
		},
		{
			input: [][]byte{
				{'S', '.', '.', '.'},
				{'.', 'A', '.', '.'},
				{'.', '.', 'M', '.'},
				{'.', '.', '.', 'X'},
			},
			start: coord{3, 3},
			want:  1,
		},
	} {
		var got int
		got += search(tc.input, 0, tc.start, N)
		got += search(tc.input, 0, tc.start, NE)
		got += search(tc.input, 0, tc.start, E)
		got += search(tc.input, 0, tc.start, SE)
		got += search(tc.input, 0, tc.start, S)
		got += search(tc.input, 0, tc.start, SW)
		got += search(tc.input, 0, tc.start, W)
		got += search(tc.input, 0, tc.start, NW)

		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}
