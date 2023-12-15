package main

import (
	"fmt"
	"testing"
)

func TestTiltNorth(t *testing.T) {
	for n, tc := range []struct {
		platform     [][]byte
		wantPlatform [][]byte
	}{
		{
			platform: [][]byte{
				[]byte("O....#...."),
				[]byte("O.OO#....#"),
				[]byte(".....##..."),
				[]byte("OO.#O....O"),
				[]byte(".O.....O#."),
				[]byte("O.#..O.#.#"),
				[]byte("..O..#O..O"),
				[]byte(".......O.."),
				[]byte("#....###.."),
				[]byte("#OO..#...."),
			},
			wantPlatform: [][]byte{
				[]byte("OOOO.#.O.."),
				[]byte("OO..#....#"),
				[]byte("OO..O##..O"),
				[]byte("O..#.OO..."),
				[]byte("........#."),
				[]byte("..#....#.#"),
				[]byte("..O..#.O.O"),
				[]byte("..O......."),
				[]byte("#....###.."),
				[]byte("#....#...."),
			},
		},
	} {
		tiltNorth(tc.platform)

		if got, want := tc.platform, tc.wantPlatform; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %s, want %s", n, got, want)
		}
	}
}

func TestTiltWest(t *testing.T) {
	for n, tc := range []struct {
		platform     [][]byte
		wantPlatform [][]byte
	}{
		{
			platform: [][]byte{
				[]byte("O....#...."),
				[]byte("O.OO#....#"),
				[]byte(".....##..."),
				[]byte("OO.#O....O"),
				[]byte(".O.....O#."),
				[]byte("O.#..O.#.#"),
				[]byte("..O..#O..O"),
				[]byte(".......O.."),
				[]byte("#....###.."),
				[]byte("#OO..#...."),
			},
			wantPlatform: [][]byte{
				[]byte("O....#...."),
				[]byte("OOO.#....#"),
				[]byte(".....##..."),
				[]byte("OO.#OO...."),
				[]byte("OO......#."),
				[]byte("O.#O...#.#"),
				[]byte("O....#OO.."),
				[]byte("O........."),
				[]byte("#....###.."),
				[]byte("#OO..#...."),
			},
		},
	} {
		tiltWest(tc.platform)

		if got, want := tc.platform, tc.wantPlatform; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %s, want %s", n, got, want)
		}
	}
}

func TestTiltSouth(t *testing.T) {
	for n, tc := range []struct {
		platform     [][]byte
		wantPlatform [][]byte
	}{
		{
			platform: [][]byte{
				[]byte("O....#...."),
				[]byte("O.OO#....#"),
				[]byte(".....##..."),
				[]byte("OO.#O....O"),
				[]byte(".O.....O#."),
				[]byte("O.#..O.#.#"),
				[]byte("..O..#O..O"),
				[]byte(".......O.."),
				[]byte("#....###.."),
				[]byte("#OO..#...."),
			},
			wantPlatform: [][]byte{
				[]byte(".....#...."),
				[]byte("....#....#"),
				[]byte("...O.##..."),
				[]byte("...#......"),
				[]byte("O.O....O#O"),
				[]byte("O.#..O.#.#"),
				[]byte("O....#...."),
				[]byte("OO....OO.."),
				[]byte("#OO..###.."),
				[]byte("#OO.O#...O"),
			},
		},
	} {
		tiltSouth(tc.platform)

		if got, want := tc.platform, tc.wantPlatform; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %s, want %s", n, got, want)
		}
	}
}

func TestTiltEast(t *testing.T) {
	for n, tc := range []struct {
		platform     [][]byte
		wantPlatform [][]byte
	}{
		{
			platform: [][]byte{
				[]byte("O....#...."),
				[]byte("O.OO#....#"),
				[]byte(".....##..."),
				[]byte("OO.#O....O"),
				[]byte(".O.....O#."),
				[]byte("O.#..O.#.#"),
				[]byte("..O..#O..O"),
				[]byte(".......O.."),
				[]byte("#....###.."),
				[]byte("#OO..#...."),
			},
			wantPlatform: [][]byte{
				[]byte("....O#...."),
				[]byte(".OOO#....#"),
				[]byte(".....##..."),
				[]byte(".OO#....OO"),
				[]byte("......OO#."),
				[]byte(".O#...O#.#"),
				[]byte("....O#..OO"),
				[]byte(".........O"),
				[]byte("#....###.."),
				[]byte("#..OO#...."),
			},
		},
	} {
		tiltEast(tc.platform)

		if got, want := tc.platform, tc.wantPlatform; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %s, want %s", n, got, want)
		}
	}
}

func TestCycle(t *testing.T) {
	for n, tc := range []struct {
		platform     [][]byte
		wantPlatform [][]byte
		cycles       int
	}{
		{
			platform: [][]byte{
				[]byte("O....#...."),
				[]byte("O.OO#....#"),
				[]byte(".....##..."),
				[]byte("OO.#O....O"),
				[]byte(".O.....O#."),
				[]byte("O.#..O.#.#"),
				[]byte("..O..#O..O"),
				[]byte(".......O.."),
				[]byte("#....###.."),
				[]byte("#OO..#...."),
			},
			wantPlatform: [][]byte{
				[]byte(".....#...."),
				[]byte("....#...O#"),
				[]byte("...OO##..."),
				[]byte(".OO#......"),
				[]byte(".....OOO#."),
				[]byte(".O#...O#.#"),
				[]byte("....O#...."),
				[]byte("......OOOO"),
				[]byte("#...O###.."),
				[]byte("#..OO#...."),
			},
			cycles: 1,
		},
		{
			platform: [][]byte{
				[]byte("O....#...."),
				[]byte("O.OO#....#"),
				[]byte(".....##..."),
				[]byte("OO.#O....O"),
				[]byte(".O.....O#."),
				[]byte("O.#..O.#.#"),
				[]byte("..O..#O..O"),
				[]byte(".......O.."),
				[]byte("#....###.."),
				[]byte("#OO..#...."),
			},
			wantPlatform: [][]byte{
				[]byte(".....#...."),
				[]byte("....#...O#"),
				[]byte(".....##..."),
				[]byte("..O#......"),
				[]byte(".....OOO#."),
				[]byte(".O#...O#.#"),
				[]byte("....O#...O"),
				[]byte(".......OOO"),
				[]byte("#..OO###.."),
				[]byte("#.OOO#...O"),
			},
			cycles: 2,
		},
		{
			platform: [][]byte{
				[]byte("O....#...."),
				[]byte("O.OO#....#"),
				[]byte(".....##..."),
				[]byte("OO.#O....O"),
				[]byte(".O.....O#."),
				[]byte("O.#..O.#.#"),
				[]byte("..O..#O..O"),
				[]byte(".......O.."),
				[]byte("#....###.."),
				[]byte("#OO..#...."),
			},
			wantPlatform: [][]byte{
				[]byte(".....#...."),
				[]byte("....#...O#"),
				[]byte(".....##..."),
				[]byte("..O#......"),
				[]byte(".....OOO#."),
				[]byte(".O#...O#.#"),
				[]byte("....O#...O"),
				[]byte(".......OOO"),
				[]byte("#...O###.O"),
				[]byte("#.OOO#...O"),
			},
			cycles: 3,
		},
	} {
		cycle(tc.platform, tc.cycles)

		if got, want := tc.platform, tc.wantPlatform; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %s, want %s", n, got, want)
		}
	}
}
