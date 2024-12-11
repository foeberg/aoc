package main

import (
	"slices"
	"testing"
)

func TestTrimLeadingZeros(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{
			input: "001",
			want:  "1",
		},
		{
			input: "00100",
			want:  "100",
		},
		{
			input: "00",
			want:  "0",
		},
		{
			input: "10",
			want:  "10",
		},
	} {
		got := trimLeadingZeros(tc.input)

		if got != tc.want {
			t.Errorf("got %s, want %s", got, tc.want)
		}
	}
}

func TestBlink(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  []string
	}{
		{
			input: "0",
			want:  []string{"1"},
		},
		{
			input: "10",
			want:  []string{"1", "0"},
		},
		{
			input: "100000",
			want:  []string{"100", "0"},
		},
		{
			input: "2",
			want:  []string{"4048"},
		},
	} {
		got := blink(tc.input)

		if !slices.Equal(got, tc.want) {
			t.Errorf("got %v\nwant %v", got, tc.want)
		}
	}
}
