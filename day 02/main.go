package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %v", err))
	}
	defer file.Close()

	powerSum := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		power := getIDPower(scanner.Text())
		powerSum += power
	}

	println(powerSum)
}

func getIDPower(line string) int {
	//Game 1: 4 blue, 7 red, 5 green; 3 blue, 4 red, 16 green; 3 red, 11 green
	posE := strings.Index(line, ":")

	blue := 0
	green := 0
	red := 0

	// 4 blue, 7 red, 5 green; 3 blue, 4 red, 16 green; 3 red, 11 green
	hands := strings.Split(line[posE+1:], ";")
	for _, hand := range hands {
		// 4 blue, 7 red, 5 green
		// 3 blue, 4 red, 16 green
		// 3 red, 11 green
		sets := strings.Split(hand, ",")
		for _, set := range sets {
			// 4 blue
			// 7 red
			// 5 green
			pair := strings.TrimSpace(set)
			parts := strings.Split(pair, " ")
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				panic("unexpected count")
			}

			switch parts[1] {
			case "blue":
				blue = max(blue, count)

			case "green":
				green = max(green, count)

			case "red":
				red = max(red, count)
			}
		}
	}

	return blue * green * red
}
