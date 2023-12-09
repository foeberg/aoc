package main

import "testing"

func TestGetType(t *testing.T) {
	for n, tc := range []struct {
		card     string
		wantType int
	}{
		{
			card:     "AAAAA",
			wantType: FiveOfKind,
		},
		{
			card:     "AABAA",
			wantType: FourOfKind,
		},
		{
			card:     "AAABB",
			wantType: FullHouse,
		},
		{
			card:     "AAA12",
			wantType: ThreeOfKind,
		},
		{
			card:     "AABBK",
			wantType: TwoPair,
		},
		{
			card:     "AAKQJ",
			wantType: OnePair,
		},
		{
			card:     "A2345",
			wantType: HighCard,
		},
	} {
		got := getType(tc.card)

		if got, want := got, tc.wantType; got != want {
			t.Errorf("[%d] got type %d, want %d", n, got, want)
		}
	}
}
