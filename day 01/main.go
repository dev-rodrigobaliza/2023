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

	total := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		total += getValue(scanner.Text())
	}

	println(total)
}

func getValue(line string) int {
	var (
		v1 int
		v2 int
	)

	for pos, char := range line {
		if char >= '0' && char <= '9' {
			v1, _ = strconv.Atoi(string(char))
			break
		}

		v := evalString(line[pos:], false)
		if v > 0 {
			v1 = v
			break
		}
	}

	for i := len(line); i >= 1; i-- {
		char := line[i-1]
		if char >= '0' && char <= '9' {
			v2, _ = strconv.Atoi(string(char))
			break
		}

		v := evalString(line[0:i], true)
		if v > 0 {
			v2 = v
			break
		}
	}

	return v1*10 + v2
}

func evalString(text string, reverse bool) int {
	var eval string

	if !reverse {
		eval = text
	} else {
		if len(text) < 3 {
			return 0
		}

		eval = text[len(text)-3:]
	}
	if pos := strings.Index(eval, "one"); pos == 0 {
		return 1
	}

	if !reverse {
		eval = text
	} else {
		if len(text) < 3 {
			return 0
		}

		eval = text[len(text)-3:]
	}
	if pos := strings.Index(eval, "two"); pos == 0 {
		return 2
	}

	if !reverse {
		eval = text
	} else {
		if len(text) < 5 {
			return 0
		}

		eval = text[len(text)-5:]
	}
	if pos := strings.Index(eval, "three"); pos == 0 {
		return 3
	}

	if !reverse {
		eval = text
	} else {
		if len(text) < 4 {
			return 0
		}

		eval = text[len(text)-4:]
	}
	if pos := strings.Index(eval, "four"); pos == 0 {
		return 4
	}

	if !reverse {
		eval = text
	} else {
		if len(text) < 4 {
			return 0
		}

		eval = text[len(text)-4:]
	}
	if pos := strings.Index(eval, "five"); pos == 0 {
		return 5
	}

	if !reverse {
		eval = text
	} else {
		if len(text) < 3 {
			return 0
		}

		eval = text[len(text)-3:]
	}
	if pos := strings.Index(eval, "six"); pos == 0 {
		return 6
	}

	if !reverse {
		eval = text
	} else {
		if len(text) < 5 {
			return 0
		}

		eval = text[len(text)-5:]
	}
	if pos := strings.Index(eval, "seven"); pos == 0 {
		return 7
	}

	if !reverse {
		eval = text
	} else {
		if len(text) < 5 {
			return 0
		}

		eval = text[len(text)-5:]
	}
	if pos := strings.Index(eval, "eight"); pos == 0 {
		return 8
	}

	if !reverse {
		eval = text
	} else {
		if len(text) < 4 {
			return 0
		}

		eval = text[len(text)-4:]
	}
	if pos := strings.Index(eval, "nine"); pos == 0 {
		return 9
	}

	return 0
}
