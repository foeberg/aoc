package main

import (
	"fmt"
	"testing"
)

func TestGetBox(t *testing.T) {
	for n, tc := range []struct {
		input    string
		wantHash int
	}{
		{
			input:    "HASH",
			wantHash: 52,
		},
		{
			input:    "rn",
			wantHash: 0,
		},
		{
			input:    "qp",
			wantHash: 1,
		},
		{
			input:    "ab",
			wantHash: 3,
		},
	} {
		gotHash := hash(tc.input)

		if got, want := gotHash, tc.wantHash; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestInsertLens(t *testing.T) {
	for n, tc := range []struct {
		boxes     map[int][]lens
		lens      lens
		wantBoxes map[int][]lens
	}{
		{
			boxes: map[int][]lens{},
			lens:  lens{"ab", 2},
			wantBoxes: map[int][]lens{
				3: {{"ab", 2}},
			},
		},
		{
			boxes: map[int][]lens{
				0: {{"rn", 1}},
			},
			lens: lens{"cm", 2},
			wantBoxes: map[int][]lens{
				0: {{"rn", 1}, {"cm", 2}},
			},
		},
		{
			boxes: map[int][]lens{
				0: {{"rn", 1}, {"cm", 2}},
			},
			lens: lens{"cm", 6},
			wantBoxes: map[int][]lens{
				0: {{"rn", 1}, {"cm", 6}},
			},
		},
		{
			boxes: map[int][]lens{
				0: {{"rn", 1}, {"cm", 2}},
			},
			lens: lens{"rn", 6},
			wantBoxes: map[int][]lens{
				0: {{"rn", 6}, {"cm", 2}},
			},
		},
	} {
		insertLens(tc.lens, tc.boxes)

		if got, want := tc.boxes, tc.wantBoxes; fmt.Sprint(tc.boxes) != fmt.Sprint(tc.wantBoxes) {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}
	}
}

func TestRemoveLens(t *testing.T) {
	for n, tc := range []struct {
		boxes     map[int][]lens
		lens      lens
		wantBoxes map[int][]lens
	}{
		{
			boxes:     map[int][]lens{},
			lens:      lens{"ab", 2},
			wantBoxes: map[int][]lens{},
		},
		{
			boxes: map[int][]lens{
				0: {{"rn", 1}},
			},
			lens: lens{"ab", 2},
			wantBoxes: map[int][]lens{
				0: {{"rn", 1}},
			},
		},
		{
			boxes: map[int][]lens{
				0: {{"rn", 1}, {"cm", 2}},
			},
			lens: lens{"rn", 0},
			wantBoxes: map[int][]lens{
				0: {{"cm", 2}},
			},
		},
		{
			boxes: map[int][]lens{
				0: {{"asd", 1}, {"rn", 1}, {"cm", 2}},
			},
			lens: lens{"rn", 0},
			wantBoxes: map[int][]lens{
				0: {{"asd", 1}, {"cm", 2}},
			},
		},
	} {
		removeLens(tc.lens, tc.boxes)

		if got, want := tc.boxes, tc.wantBoxes; fmt.Sprint(tc.boxes) != fmt.Sprint(tc.wantBoxes) {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}
	}
}

func TestParseLens(t *testing.T) {
	for n, tc := range []struct {
		input    string
		wantLens lens
	}{
		{
			input:    "ab=2",
			wantLens: lens{"ab", 2},
		},
		{
			input:    "ab-",
			wantLens: lens{"ab", 0},
		},
	} {
		gotLens := parseLens(tc.input)

		if got, want := gotLens, tc.wantLens; got != want {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}
	}
}
