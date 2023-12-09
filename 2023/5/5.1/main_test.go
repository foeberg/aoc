package main

import "testing"

func TestSourceToDest(t *testing.T) {
	for _, tc := range []struct {
		source          int
		mappings        []mapping
		wantDestination int
	}{
		{
			source: 79,
			mappings: []mapping{
				{destination: 50, source: 98, len: 2},
				{destination: 52, source: 50, len: 48},
			},
			wantDestination: 81,
		},
	} {
		gotDestination := sourceToDest(tc.source, tc.mappings)

		if got, want := gotDestination, tc.wantDestination; got != want {
			t.Errorf("got destination %d, want %d", got, want)
		}
	}
}
