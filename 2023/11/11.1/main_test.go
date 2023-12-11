package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestDistance(t *testing.T) {
	for n, tc := range []struct {
		in   pair
		want int
	}{
		{
			in: pair{
				coord{0, 2},
				coord{12, 7},
			},
			want: 17,
		},
		{
			in: pair{
				coord{12, 7},
				coord{0, 2},
			},
			want: 17,
		},
		{
			in: pair{
				coord{4, 0},
				coord{9, 10},
			},
			want: 15,
		},
		{
			in: pair{
				coord{9, 10},
				coord{4, 0},
			},
			want: 15,
		},
		{
			in: pair{
				coord{0, 11},
				coord{5, 11},
			},
			want: 5,
		},
		{
			in: pair{
				coord{5, 11},
				coord{0, 11},
			},
			want: 5,
		},
	} {
		got := distance(tc.in)

		if got, want := got, tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestCreatePairs(t *testing.T) {
	for n, tc := range []struct {
		in   []coord
		want []pair
	}{
		{
			in: []coord{
				{0, 0},
				{1, 1},
				{2, 2},
				{3, 3},
			},
			want: []pair{
				{coord{0, 0}, coord{1, 1}},
				{coord{0, 0}, coord{2, 2}},
				{coord{0, 0}, coord{3, 3}},
				{coord{1, 1}, coord{2, 2}},
				{coord{1, 1}, coord{3, 3}},
				{coord{2, 2}, coord{3, 3}},
			},
		},
	} {
		got := createPairs(tc.in)

		if got, want := got, tc.want; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}
	}
}

func TestExpandImage(t *testing.T) {
	for n, tc := range []struct {
		in   []string
		want []string
	}{
		{
			in: []string{
				".....",
				"..#..",
				".....",
				"#...#",
			},
			want: []string{
				".......",
				".......",
				"...#...",
				".......",
				".......",
				"#.....#",
			},
		},
	} {
		got := expandImage(tc.in)

		if got, want := got, tc.want; fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("[%d] got %s, want %s", n, got, want)
		}
	}
}

func TestFindGalaxies(t *testing.T) {
	for n, tc := range []struct {
		image []string
		want  []coord
	}{
		{
			image: []string{
				".#...",
				".....",
				"...#.",
				"#...#",
				".#...",
			},
			want: []coord{
				{1, 0},
				{3, 2},
				{0, 3},
				{4, 3},
				{1, 4},
			},
		},
	} {
		got := findGalaxies(tc.image)

		if got, want := got, tc.want; !slices.Equal(got, want) {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}
	}
}
