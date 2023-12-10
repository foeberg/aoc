package main

import (
	"testing"
)

func TestLcm(t *testing.T) {
	for n, tc := range []struct {
		nums []int
		want int
	}{
		{
			nums: []int{30, 51},
			want: 510,
		},
		{
			nums: []int{271, 51, 966},
			want: 4450362,
		},
	} {
		got := lcm(tc.nums...)

		if got, want := got, tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
