package main

import (
	"fmt"
	"testing"
)

func TestReduceRanges(t *testing.T) {
	for _, tc := range []struct {
		desc       string
		input      [][]int
		wantOutput [][]int
	}{
		{
			desc:       "not overlapping",
			input:      [][]int{{2, 5}, {6, 8}},
			wantOutput: [][]int{{2, 5}, {6, 8}},
		},
		{
			desc:       "range contained within",
			input:      [][]int{{2, 10}, {3, 5}},
			wantOutput: [][]int{{2, 10}},
		},
		{
			desc:       "overlapping",
			input:      [][]int{{3, 9}, {2, 5}},
			wantOutput: [][]int{{2, 9}},
		},
	} {
		gotOutput := reduceRanges(tc.input)

		if got, want := gotOutput, tc.wantOutput; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got output\n%v, want\n%v", got, want)
		}
	}
}
