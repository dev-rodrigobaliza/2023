package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type hand struct {
	cards  string
	points int
	bid    int
}

func main() {
	lines := strings.Split(strings.TrimSpace(input), "\r\n")

	hands := getHands(lines)
	winnings := getWinnings(hands)
	println(winnings)

	hands2 := getHands2(lines)
	winnings2 := getWinnings2(hands2)
	println(winnings2)

}

func getWinnings2(hands []hand) int {
	slices.SortFunc(hands, func(a, b hand) int {
		if a.points != b.points {
			return a.points - b.points
		}

		chks := []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
		for i := range a.cards {
			if a.cards[i] == b.cards[i] {
				continue
			}
			for _, chk := range chks {
				if a.cards[i] == chk {
					return 1
				}
				if b.cards[i] == chk {
					return -1
				}
			}

			panic("unknow card to compare...")
		}

		panic("both cards are the same...")
	})

	w := 0
	for i, hand := range hands {
		fmt.Printf("{Cards:%s Wager:%d Strenght:%d}\n", hand.cards, hand.bid, hand.points-1)
		w += (i + 1) * hand.bid
	}

	return w
}

func getHands2(lines []string) []hand {
	hands := make([]hand, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		points := getPoints2(parts[0])
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, hand{
			cards:  parts[0],
			points: points,
			bid:    bid,
		})
	}

	return hands
}

func getPoints2(cards string) int {
	have := []rune{}
	diff := []rune{}
	jokers := 0

	cs := []rune(cards)
	slices.Sort(cs)
	for _, c := range cs {
		if c == 'J' {
			jokers++
			continue
		}
		if !slices.Contains(have, c) {
			have = append(have, c)
		} else {
			if !slices.Contains(diff, c) {
				diff = append(diff, c)
			}
		}
	}

	switch jokers {
	case 5:
		// five of a kind
		return 7

	case 4:
		// five of a kind
		return 7

	case 3:
		if len(have) == 1 {
			// five of a kind
			return 7
		}

		// four of a kind
		return 6

	case 2:
		switch len(have) {
		case 1:
			// five of a kind
			return 7

		case 2:
			// four of a kind
			return 6

		default:
			// three of a kind
			return 4
		}

	case 1:
		switch len(have) {
		case 1:
			// five of a kind
			return 7

		case 2:
			if len(diff) == 1 {
				// four of a kind
				return 6
			}

			// full house
			return 5

		case 3:
			// three of a kind
			return 4

		case 4:
			// one pair
			return 2
		}
	}

	switch len(have) {
	case 5:
		// high card
		return 1

	case 4:
		// one pair
		return 2

	case 3:
		if len(diff) == 2 {
			// two pair
			return 3
		}

		// three of a kind
		return 4

	case 2:
		if len(diff) == 2 {
			// full house
			return 5
		}

		// four of a kind
		return 6

	case 1:
		// five of a kind
		return 7
	}

	panic("houston, we have a problem here...")
}

func getWinnings(hands []hand) int {
	slices.SortFunc(hands, func(a, b hand) int {
		if a.points != b.points {
			return a.points - b.points
		}

		chks := []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
		for i := range a.cards {
			if a.cards[i] == b.cards[i] {
				continue
			}
			for _, chk := range chks {
				if a.cards[i] == chk {
					return 1
				} else if b.cards[i] == chk {
					return -1
				}
			}

			panic("unknow card to compare...")
		}

		panic("both cards are the same...")
	})

	w := 0
	for i, hand := range hands {
		w += (i + 1) * hand.bid
	}

	return w
}

func getHands(lines []string) []hand {
	hands := make([]hand, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		points := getPoints(parts[0])
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, hand{
			cards:  parts[0],
			points: points,
			bid:    bid,
		})
	}

	return hands
}

func getPoints(cards string) int {
	have := []rune{}
	diff := []rune{}

	cs := []rune(cards)
	slices.Sort(cs)
	for _, c := range cs {
		if !slices.Contains(have, c) {
			have = append(have, c)
		} else {
			if !slices.Contains(diff, c) {
				diff = append(diff, c)
			}
		}
	}

	switch len(have) {
	case 5:
		// high card
		return 1

	case 4:
		// one pair
		return 2

	case 3:
		if len(diff) == 2 {
			// two pair
			return 3
		}

		// three of a kind
		return 4

	case 2:
		if len(diff) == 2 {
			// full house
			return 5
		}

		// four of a kind
		return 6

	case 1:
		// five of a kind
		return 7

	default:
		panic("houston, we have a problem here...")
	}
}
