package main

import (
	"slices"
	"testing"
)

func TestFindHorizontalDuplicates(t *testing.T) {
	for n, tc := range []struct {
		pattern  []string
		wantRows []duplicate
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
			wantRows: []duplicate{{0, true}, {3, false}},
		},
		{
			pattern: []string{
				"#...##..#",
				"#...###.#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			wantRows: []duplicate{{0, true}, {3, false}},
		},
	} {
		gotRows := findHorizontalDuplicates(tc.pattern)

		if got, want := gotRows, tc.wantRows; !slices.Equal(got, want) {
			t.Errorf("[%d] got duplicates %v, want %v", n, got, want)
		}
	}
}

func TestFindVerticalDuplicates(t *testing.T) {
	for n, tc := range []struct {
		pattern  []string
		wantRows []duplicate
	}{
		{
			pattern: []string{
				"..##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			wantRows: []duplicate{{0, true}, {4, false}},
		},
		{
			pattern: []string{
				"..##..##..",
				"..#.##.#..",
				"##......##",
				".#......##",
				"..#.##.#..",
				"..##..##..",
				"..#.##.#..",
			},
			wantRows: []duplicate{{0, true}, {4, false}, {8, false}},
		},
	} {
		gotRows := findVerticalDuplicates(tc.pattern)

		if got, want := gotRows, tc.wantRows; !slices.Equal(got, want) {
			t.Errorf("[%d] got duplicatess %v, want %v", n, got, want)
		}
	}
}

func TestSummarizePattern(t *testing.T) {
	for n, tc := range []struct {
		pattern []string
		wantNum int
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
			wantNum: 300,
		},
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
			wantNum: 100,
		},
	} {
		gotNum := summarizePattern(tc.pattern)

		if got, want := gotNum, tc.wantNum; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
