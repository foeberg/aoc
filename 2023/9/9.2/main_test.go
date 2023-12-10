package main

import (
	"slices"
	"testing"
)

func TestExtrapolate(t *testing.T) {
	for n, tc := range []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{2, 4, 6, 8},
			want: []int{0, 2, 4, 6, 8},
		},
		{
			in:   []int{-2, -4, -6, -8},
			want: []int{0, -2, -4, -6, -8},
		},
		{
			in:   []int{0, 0, 0},
			want: []int{0, 0, 0, 0},
		},
		{
			in:   []int{1, 1, 1, 1},
			want: []int{1, 1, 1, 1, 1},
		},
	} {
		got := extrapolate(tc.in)

		if got, want := got, tc.want; !slices.Equal(got, want) {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
