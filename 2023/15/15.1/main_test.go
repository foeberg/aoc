package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	for n, tc := range []struct {
		input    string
		wantHash int
	}{
		{
			input:    "HASH",
			wantHash: 52,
		},
	} {
		gotHash := hash(tc.input)

		if got, want := gotHash, tc.wantHash; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
