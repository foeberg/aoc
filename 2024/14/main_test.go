package main

import "testing"

func TestFindCycle(t *testing.T) {
	for _, tc := range []struct {
		sizeX     int
		sizeY     int
		robot     robot
		wantCycle int
	}{
		{
			sizeX:     10,
			sizeY:     5,
			robot:     robot{0, 0, 0, 1},
			wantCycle: 5,
		},
		{
			sizeX:     10,
			sizeY:     5,
			robot:     robot{0, 0, 1, 0},
			wantCycle: 10,
		},
		{
			sizeX:     10,
			sizeY:     5,
			robot:     robot{0, 0, 0, 2},
			wantCycle: 5,
		},
		{
			sizeX:     10,
			sizeY:     5,
			robot:     robot{0, 0, 0, 3},
			wantCycle: 5,
		},
		{
			sizeX:     4,
			sizeY:     4,
			robot:     robot{0, 0, 1, 1},
			wantCycle: 4,
		},
		{
			sizeX:     4,
			sizeY:     4,
			robot:     robot{0, 0, 1, -1},
			wantCycle: 4,
		},
		{
			sizeX:     4,
			sizeY:     5,
			robot:     robot{0, 0, -1, -1},
			wantCycle: 20,
		},
	} {
		gotCycle := findCycle(tc.sizeX, tc.sizeY, tc.robot)

		if got, want := gotCycle, tc.wantCycle; got != want {
			t.Errorf("got cycle %d, want %d", got, want)
		}
	}
}

func TestSimulate(t *testing.T) {
	for _, tc := range []struct {
		sizeX   int
		sizeY   int
		seconds int
		robot   robot
		wantX   int
		wantY   int
	}{
		{
			sizeX:   4,
			sizeY:   4,
			seconds: 5,
			robot:   robot{0, 0, 1, 0},
			wantX:   1,
			wantY:   0,
		},
		{
			sizeX:   4,
			sizeY:   4,
			seconds: 5,
			robot:   robot{0, 0, 2, 0},
			wantX:   2,
			wantY:   0,
		},
		{
			sizeX:   4,
			sizeY:   4,
			seconds: 5,
			robot:   robot{0, 0, 0, 1},
			wantX:   0,
			wantY:   1,
		},
		{
			sizeX:   4,
			sizeY:   4,
			seconds: 4,
			robot:   robot{0, 0, 3, -1},
			wantX:   0,
			wantY:   0,
		},
		{
			sizeX:   7,
			sizeY:   11,
			seconds: 2,
			robot:   robot{0, 4, 3, -3},
			wantX:   6,
			wantY:   9,
		},
	} {
		gotX, gotY := simulate(tc.sizeX, tc.sizeY, tc.seconds, tc.robot)

		if gotX != tc.wantX {
			t.Errorf("got X %d, want %d", gotX, tc.wantX)
		}

		if gotY != tc.wantY {
			t.Errorf("got Y %d, want %d", gotY, tc.wantY)
		}

	}
}
