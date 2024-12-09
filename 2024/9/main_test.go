package main

import "testing"

func TestFindLastFile(t *testing.T) {
	for n, tc := range []struct {
		input     []int
		wantSize  int
		wantIndex int
	}{
		{
			input:     []int{0, 0, -1, -1, 1, -1, -1, -1, 2, 2, 2},
			wantSize:  3,
			wantIndex: 8,
		},
		{
			input:     []int{-1, -1, -1},
			wantSize:  0,
			wantIndex: 0,
		},
		{
			input:     []int{1, 1, 1},
			wantSize:  3,
			wantIndex: 0,
		},
		{
			input:     []int{1, 1, 1, -1, -1, -1},
			wantSize:  3,
			wantIndex: 0,
		},
		{
			input:     []int{-1, 1, 1, 1, -1, -1, -1},
			wantSize:  3,
			wantIndex: 1,
		},
	} {
		gotIndex, gotSize := findLastFile(tc.input)

		if got, want := gotIndex, tc.wantIndex; got != want {
			t.Errorf("[%d] got index: %d, want %d", n, got, want)
		}
		if got, want := gotSize, tc.wantSize; got != want {
			t.Errorf("[%d] got size: %d, want %d", n, got, want)
		}
	}
}

func TestFindGap(t *testing.T) {
	for n, tc := range []struct {
		input     []int
		size      int
		wantIndex int
		wantOK    bool
	}{
		{
			input:     []int{0, 0, -1, -1, 1},
			size:      2,
			wantIndex: 2,
			wantOK:    true,
		},
		{
			input:     []int{0, 0, -1, -1, 1},
			size:      1,
			wantIndex: 2,
			wantOK:    true,
		},
		{
			input:     []int{-1, -1, 1},
			size:      1,
			wantIndex: 0,
			wantOK:    true,
		},
		{
			input: []int{0, 0, -1, -1, 1},
			size:  3,
		},
	} {
		gotIndex, gotOK := findGap(tc.input, tc.size)

		if got, want := gotIndex, tc.wantIndex; got != want {
			t.Errorf("[%d] got index: %d, want %d", n, got, want)
		}
		if got, want := gotOK, tc.wantOK; got != want {
			t.Errorf("[%d] got OK: %v, want %v", n, got, want)
		}
	}
}
