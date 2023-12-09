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
			card:     "AAA23",
			wantType: ThreeOfKind,
		},
		{
			card:     "AABBK",
			wantType: TwoPair,
		},
		{
			card:     "AAKQJ",
			wantType: ThreeOfKind,
		},
		{
			card:     "A2345",
			wantType: HighCard,
		},
		{
			card:     "AAAAJ",
			wantType: FiveOfKind,
		},
		{
			card:     "AAJJJ",
			wantType: FiveOfKind,
		},
		{
			card:     "AAAJJ",
			wantType: FiveOfKind,
		},
		{
			card:     "JJJJ2",
			wantType: FiveOfKind,
		},
		{
			card:     "AAJJ2",
			wantType: FourOfKind,
		},
		{
			card:     "AAA2J",
			wantType: FourOfKind,
		},
		{
			card:     "JJJ23",
			wantType: FourOfKind,
		},
		{
			card:     "AAJ23",
			wantType: ThreeOfKind,
		},
		{
			card:     "JJ234",
			wantType: ThreeOfKind,
		},
		{
			card:     "AA22J",
			wantType: FullHouse,
		},
		{
			card:     "A123J",
			wantType: OnePair,
		},
		{
			card:     "A123J",
			wantType: OnePair,
		},
	} {
		got := getType(tc.card)

		if got, want := got, tc.wantType; got != want {
			t.Errorf("[%d] got type %d, want %d", n, got, want)
		}
	}
}
