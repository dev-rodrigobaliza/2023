package main

import (
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type pos struct {
	line int
	col  int
}

type number struct {
	line     int
	colStart int
	colEnd   int
	value    int
}

func main() {
	lines := strings.Split(input, "\r\n")
	sum := getNumbersSum(lines)
	println(sum)

	numbers := getNumbers(lines)
	sum = 0
	for _, ns := range numbers {
		if len(ns) == 2 {
			sum += ns[0].value * ns[1].value
		}
	}
	println(sum)
}

func getNumbers(lines []string) map[pos][]number {
	numbers := make(map[pos][]number, 0)

	maybeNumber := false
	posS := 0

	for l, line := range lines {
		for c, char := range line {
			if char >= '0' && char <= '9' {
				if !maybeNumber {
					maybeNumber = true
					posS = c
				} else if c == len(line)-1 {
					p := isNumber(lines, l, posS, c, l == 0, l == len(lines)-1)
					if p != nil {
						ns, ok := numbers[*p]
						if !ok {
							ns = make([]number, 0)
						}

						n, _ := strconv.Atoi(line[posS : c+1])
						ns = append(ns, number{
							line:     l,
							colStart: posS,
							colEnd:   c,
							value:    n,
						})

						numbers[*p] = ns
					}

					maybeNumber = false
				}
			} else {
				if maybeNumber {
					p := isNumber(lines, l, posS, c, l == 0, l == len(lines)-1)
					if p != nil {
						ns, ok := numbers[*p]
						if !ok {
							ns = make([]number, 0)
						}

						n, _ := strconv.Atoi(line[posS:c])
						ns = append(ns, number{
							line:     l,
							colStart: posS,
							colEnd:   c,
							value:    n,
						})

						numbers[*p] = ns
					}

					maybeNumber = false
				}
			}
		}
	}

	return numbers
}

func isNumber(lines []string, ver, horS, horE int, first, last bool) *pos {
	cS := horS - 1
	if cS < 0 {
		cS = 0
	}
	cE := horE + 1
	if cE > len(lines[0]) {
		cE = len(lines[0])
	}

	for l := ver - 1; l <= ver+1; l++ {
		if (first && l == ver-1) || (last && l == ver+1) {
			continue
		}
		check := lines[l][cS:cE]
		for c, char := range check {
			if char == '*' {
				p := pos{
					line: l,
					col:  cS + c,
				}
				return &p
			}
		}
	}

	return nil
}

func getNumbersSum(lines []string) int {
	sum := 0

	maybeNumber := false
	posS := 0
	for l, line := range lines {
		for c, char := range line {
			if char >= '0' && char <= '9' {
				if !maybeNumber {
					maybeNumber = true
					posS = c
				} else if c == len(line)-1 {
					if isPartNumber(lines, l, posS, c, l == 0, l == len(lines)-1) {
						n, _ := strconv.Atoi(line[posS : c+1])
						sum += n
					}
					maybeNumber = false
				}
			} else {
				if maybeNumber {
					if isPartNumber(lines, l, posS, c, l == 0, l == len(lines)-1) {
						n, _ := strconv.Atoi(line[posS:c])
						sum += n
					}
					maybeNumber = false
				}
			}
		}
	}

	return sum
}

func isPartNumber(lines []string, ver, horS, horE int, first, last bool) bool {
	cS := horS - 1
	if cS < 0 {
		cS = 0
	}
	cE := horE + 1
	if cE > len(lines[0]) {
		cE = len(lines[0])
	}

	for l := ver - 1; l <= ver+1; l++ {
		if (first && l == ver-1) || (last && l == ver+1) {
			continue
		}
		check := lines[l][cS:cE]
		for _, char := range check {
			if char != '.' && (char < '0' || char > '9') {
				return true
			}
		}
	}

	return false
}
