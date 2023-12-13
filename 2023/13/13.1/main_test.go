package main

import (
	"slices"
	"testing"
)

func TestFindHorizontalDuplicates(t *testing.T) {
	for n, tc := range []struct {
		pattern  []string
		wantRows []int
	}{
		{
			pattern: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			wantRows: []int{3},
		},
		{
			pattern: []string{
				"#...##..#",
				"#...##..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			wantRows: []int{0, 3},
		},
	} {
		gotRows := findHorizontalDuplicates(tc.pattern)

		if got, want := gotRows, tc.wantRows; !slices.Equal(got, want) {
			t.Errorf("[%d] got duplicates %d, want %d", n, got, want)
		}
	}
}

func TestFindVerticalDuplicates(t *testing.T) {
	for n, tc := range []struct {
		pattern  []string
		wantRows []int
	}{
		{
			pattern: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			wantRows: []int{4},
		},
		{
			pattern: []string{
				"#.##..##..",
				"..#.##.#..",
				"##......##",
				"##......##",
				"..#.##.#..",
				"..##..##..",
				"#.#.##.#..",
			},
			wantRows: []int{4, 8},
		},
	} {
		gotRows := findVerticalDuplicates(tc.pattern)

		if got, want := gotRows, tc.wantRows; !slices.Equal(got, want) {
			t.Errorf("[%d] got duplicates %d, want %d", n, got, want)
		}
	}
}

func TestSummarizePattern(t *testing.T) {
	for n, tc := range []struct {
		pattern []string
		wantNum int
	}{
		{
			// No reflections
			pattern: []string{
				"#.##..##.",
				"..#.##.#.",
				"###.....#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			wantNum: 0,
		},
		{
			// Horizontal reflection
			pattern: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			wantNum: 400,
		},
		{
			// Vertical reflection
			pattern: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			wantNum: 5,
		},
	} {
		gotNum := summarizePattern(tc.pattern)

		if got, want := gotNum, tc.wantNum; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
