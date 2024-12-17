package main

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	for n, tc := range []struct {
		input      [][]byte
		position   coord
		direction  byte
		wantNewPos coord
		wantOutput [][]byte
	}{
		// Right
		{
			input: [][]byte{
				{'@', 'O', '.', '#'},
			},
			position:   coord{0, 0},
			direction:  '>',
			wantNewPos: coord{1, 0},
			wantOutput: [][]byte{
				{'.', '@', 'O', '#'},
			},
		},
		{
			input: [][]byte{
				{'@', 'O', 'O', '.', '#'},
			},
			position:   coord{0, 0},
			direction:  '>',
			wantNewPos: coord{1, 0},
			wantOutput: [][]byte{
				{'.', '@', 'O', 'O', '#'},
			},
		},
		{
			input: [][]byte{
				{'@', 'O', '#'},
			},
			position:   coord{1, 0},
			direction:  '>',
			wantNewPos: coord{1, 0},
			wantOutput: [][]byte{
				{'@', 'O', '#'},
			},
		},
		// Left
		{
			input: [][]byte{
				{'#', '.', 'O', '@'},
			},
			position:   coord{3, 0},
			direction:  '<',
			wantNewPos: coord{2, 0},
			wantOutput: [][]byte{
				{'#', 'O', '@', '.'},
			},
		},
		{
			input: [][]byte{
				{'#', '.', 'O', 'O', '@'},
			},
			position:   coord{4, 0},
			direction:  '<',
			wantNewPos: coord{3, 0},
			wantOutput: [][]byte{
				{'#', 'O', 'O', '@', '.'},
			},
		},
		{
			input: [][]byte{
				{'#', 'O', '@'},
			},
			position:   coord{2, 0},
			direction:  '<',
			wantNewPos: coord{2, 0},
			wantOutput: [][]byte{
				{'#', 'O', '@'},
			},
		},
		// Up
		{
			input: [][]byte{
				{'#'},
				{'.'},
				{'O'},
				{'@'},
			},
			position:   coord{0, 3},
			direction:  '^',
			wantNewPos: coord{0, 2},
			wantOutput: [][]byte{
				{'#'},
				{'O'},
				{'@'},
				{'.'},
			},
		},
		{
			input: [][]byte{
				{'#'},
				{'.'},
				{'O'},
				{'O'},
				{'@'},
			},
			position:   coord{0, 4},
			direction:  '^',
			wantNewPos: coord{0, 3},
			wantOutput: [][]byte{
				{'#'},
				{'O'},
				{'O'},
				{'@'},
				{'.'},
			},
		},
		{
			input: [][]byte{
				{'#'},
				{'O'},
				{'@'},
			},
			position:   coord{0, 2},
			direction:  '^',
			wantNewPos: coord{0, 2},
			wantOutput: [][]byte{
				{'#'},
				{'O'},
				{'@'},
			},
		},
		// Down
		{
			input: [][]byte{
				{'@'},
				{'O'},
				{'.'},
				{'#'},
			},
			position:   coord{0, 0},
			direction:  'v',
			wantNewPos: coord{0, 1},
			wantOutput: [][]byte{
				{'.'},
				{'@'},
				{'O'},
				{'#'},
			},
		},
		{
			input: [][]byte{
				{'@'},
				{'O'},
				{'O'},
				{'.'},
				{'#'},
			},
			position:   coord{0, 0},
			direction:  'v',
			wantNewPos: coord{0, 1},
			wantOutput: [][]byte{
				{'.'},
				{'@'},
				{'O'},
				{'O'},
				{'#'},
			},
		},
		{
			input: [][]byte{
				{'@'},
				{'O'},
				{'#'},
			},
			position:   coord{0, 0},
			direction:  'v',
			wantNewPos: coord{0, 0},
			wantOutput: [][]byte{
				{'@'},
				{'O'},
				{'#'},
			},
		},
		// Part 2
		{
			input: [][]byte{
				[]byte("....@..."),
				[]byte("...[]..."),
				[]byte("..[][].."),
				[]byte("..[][].."),
				[]byte("..[][].."),
				[]byte("..##...."),
			},
			position:   coord{4, 0},
			direction:  'v',
			wantNewPos: coord{4, 0},
			wantOutput: [][]byte{
				[]byte("....@..."),
				[]byte("...[]..."),
				[]byte("..[][].."),
				[]byte("..[][].."),
				[]byte("..[][].."),
				[]byte("..##...."),
			},
		},
	} {
		gotNewPos := move(tc.input, tc.position, tc.direction)

		if got, want := gotNewPos, tc.wantNewPos; gotNewPos != tc.wantNewPos {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}

		if got, want := tc.input, tc.wantOutput; fmt.Sprint(got) != fmt.Sprint(want) {
			draw(got)
			fmt.Println()
			draw(want)
			fmt.Println()
			t.Errorf("[%d] unexpected output", n)
		}
	}
}
