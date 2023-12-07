package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type race struct {
	time     int
	distance int
	ways     int
}

func (r *race) setWays() {
	for i := 0; i <= r.time; i++ {
		d := (r.time-i)*i
		if d > r.distance {
			r.ways++
		}
	}
}

func main() {
	races := getRaces()
	ways := 1
	for _, r := range races {
		r.setWays()
		ways *= r.ways
	}

	println(ways)

	races2 := getRaces2()
	ways2 := 1
	for _, r := range races2 {
		r.setWays()
		ways2 *= r.ways
	}

	println(ways2)
}

func getRaces2() []*race {
	races := make([]*race, 0)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	times := getNumbers(strings.Replace(lines[0], " ", "", -1))
	distances := getNumbers(strings.Replace(lines[1], " ", "", -1))

	for i := 0; i < len(times); i++ {
		races = append(races, &race{
			time:     times[i],
			distance: distances[i],
		})
	}

	return races
}

func getRaces() []*race {
	races := make([]*race, 0)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	times := getNumbers(strings.TrimSpace(lines[0]))
	distances := getNumbers(strings.TrimSpace(lines[1]))

	for i := 0; i < len(times); i++ {
		races = append(races, &race{
			time:     times[i],
			distance: distances[i],
		})
	}

	return races
}

func getNumbers(text string) []int {
	numbers := []int{}

	re := regexp.MustCompile(`\d+`)
	parts := strings.Split(text, ":")[1]
	for _, val := range re.FindAllStringSubmatch(parts, -1) {
		num, _ := strconv.Atoi(val[0])
		numbers = append(numbers, num)
	}

	return numbers
}
