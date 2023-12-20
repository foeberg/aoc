package main

import (
	"testing"
)

func TestCalculateArea(t *testing.T) {
	for n, tc := range []struct {
		vertices []coord
		wantArea int
	}{
		{
			vertices: []coord{
				{0, 0},
				{1, 0},
				{1, 1},
				{0, 1},
			},
			wantArea: 4,
		},
		{
			vertices: []coord{
				{0, 0},
				{2, 0},
				{2, 2},
				{0, 2},
			},
			wantArea: 9,
		},
	} {
		gotArea := calculateArea(tc.vertices)

		if got, want := gotArea, tc.wantArea; got != want {
			t.Errorf("[%d] got area %d, want %d", n, got, want)
		}
	}
}
