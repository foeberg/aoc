package main

import "testing"

func TestCalculateJoltage(t *testing.T) {
	for _, tc := range []struct {
		bank        []int
		wantJoltage int
	}{
		{
			bank:        []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			wantJoltage: 98,
		},
		{
			bank:        []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			wantJoltage: 89,
		},
		{
			bank:        []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			wantJoltage: 78,
		},
		{
			bank:        []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			wantJoltage: 92,
		},
	} {
		gotJoltage := calculateJoltage(tc.bank)

		if got, want := gotJoltage, tc.wantJoltage; got != want {
			t.Errorf("got joltage %d, want %d", got, want)
		}
	}
}

func TestCalculateJoltage2(t *testing.T) {
	for _, tc := range []struct {
		bank        []int
		wantJoltage int
	}{
		{
			bank:        []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			wantJoltage: 987654321111,
		},
		{
			bank:        []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			wantJoltage: 811111111119,
		},
		{
			bank:        []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			wantJoltage: 434234234278,
		},
		{
			bank:        []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			wantJoltage: 888911112111,
		},
	} {
		gotJoltage := calculateJoltage2(tc.bank, 12)

		if got, want := gotJoltage, tc.wantJoltage; got != want {
			t.Errorf("got joltage %d, want %d", got, want)
		}
	}
}

func TestPow(t *testing.T) {
	for _, tc := range []struct {
		base    int
		exp     int
		wantRes int
	}{
		{
			base:    10,
			exp:     3,
			wantRes: 1000,
		},
		{
			base:    10,
			exp:     5,
			wantRes: 100000,
		},
		{
			base:    10,
			exp:     6,
			wantRes: 1000000,
		},
	} {
		if got, want := pow(tc.base, tc.exp), tc.wantRes; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
