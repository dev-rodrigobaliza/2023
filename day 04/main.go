package main

import (
	"slices"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	cards := strings.Split(input, "\r\n")

	winners, games := getGames(cards)
	sum := getPoints(winners, games)
	println(sum)

	sum = getCards(winners, games)
	println(sum)
}

func getCards(winners map[int][]int, games map[int][]int) int {
	sum := 0
	copies := make(map[int]int, len(winners))

	for i := 1; i <= len(winners); i++ {
		copies[i] = copies[i] + 1

		ws := winners[i]
		gs := games[i]
		ps := calcWins(gs, ws)

		for m := 1; m <= copies[i]; m++ {
			for j := 1; j <= ps; j++ {
				k := i+j
				wc := copies[k]
				copies[k] = wc+1
			}
		}

		sum += copies[i]
	}

	return sum
}

func calcWins(winners []int, games []int) int {
	points := 0

	for _, g := range games {
		if slices.Contains(winners, g) {
			points++
		}
	}

	return points
}

func getPoints(winners map[int][]int, games map[int][]int) int {
	sum := 0

	for card, gs := range games {
		ws := winners[card]

		sum += calcPoints(gs, ws)
	}

	return sum
}

func calcPoints(winners []int, games []int) int {
	points := 0

	for _, g := range games {
		if slices.Contains(winners, g) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return points
}

func getGames(cards []string) (map[int][]int, map[int][]int) {
	winners := make(map[int][]int, len(cards))
	games := make(map[int][]int, len(cards))

	for _, card := range cards {
		wns := make([]int, 0)
		gns := make([]int, 0)

		pos := strings.Index(card, ":")
		parts := strings.Split(card[:pos], " ")
		p := 1
		for {
			if parts[p] != "" {
				break
			}

			p++
		}
		n, _ := strconv.Atoi(parts[p])

		parts = strings.Split(card[pos+1:], "|")

		ws := strings.Split(parts[0], " ")
		for _, w := range ws {
			n, _ := strconv.Atoi(strings.TrimSpace(w))
			if n > 0 {
				wns = append(wns, n)
			}
		}
		winners[n] = wns

		gs := strings.Split(parts[1], " ")
		for _, g := range gs {
			n, _ := strconv.Atoi(strings.TrimSpace(g))
			if n > 0 {
				gns = append(gns, n)
			}
		}
		games[n] = gns
	}

	return winners, games
}
