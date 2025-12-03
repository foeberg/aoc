package main

import "testing"

func TestFindDoubleSequence(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want int
	}{
		{
			num:  "101",
			want: 0,
		},
		{
			num:  "11",
			want: 11,
		},
		{
			num:  "1010",
			want: 1010,
		},
		{
			num:  "1188511885",
			want: 1188511885,
		},
	} {
		got := findDoubleSequence(tc.num)
		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}

func TestFindSequences(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want int
	}{
		{
			num:  "101",
			want: 0,
		},
		{
			num:  "11",
			want: 11,
		},
		{
			num:  "1010",
			want: 1010,
		},
		{
			num:  "1188511885",
			want: 1188511885,
		},
		{
			num:  "999",
			want: 999,
		},
		{
			num:  "565656",
			want: 565656,
		},
		{
			num:  "824824824",
			want: 824824824,
		},
		{
			num:  "2121212121",
			want: 2121212121,
		},
		{
			num:  "999969",
			want: 0,
		},
		{
			num:  "6666668966",
			want: 0,
		},
	} {
		got := findSequences(tc.num)
		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}
