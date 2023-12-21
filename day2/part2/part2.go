package part2

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day 1 Part 2")
	ex, _ := os.Executable()
	f := path.Join(filepath.Dir(ex), "day1", "./input.txt")

	data, err := os.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(SumOfLines(string(data)))
}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}
func GetFirstAndLastDigits(s string) (string, string) {
	var numbers = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	re, _ := regexp.Compile(fmt.Sprintf("([0-9]|%s)", strings.Join(numbers[:], "|")))

	var firstMatch string
	var lastMatch string

	for i := 0; i < len(s); i++ {
		match := re.FindString(s[i:])
		if match == "" {
			continue
		}
		if firstMatch == "" {
			firstMatch = match
		}
		lastMatch = match
	}

	isFirstNumber, _ := regexp.MatchString("[0-9]", firstMatch)
	if !isFirstNumber {
		firstMatch = fmt.Sprint(Find(numbers[:], firstMatch) + 1)
	}
	isSecondNumber, _ := regexp.MatchString("[0-9]", lastMatch)
	if !isSecondNumber {
		lastMatch = fmt.Sprint(Find(numbers[:], lastMatch) + 1)
	}
	return firstMatch, lastMatch
}

func GetConcatenatedDigits(s string) (int, error) {
	first, last := GetFirstAndLastDigits(s)
	return strconv.Atoi(first + last)
}

func SumOfLines(s string) (int, error) {
	lines := strings.Split(s, "\n")
	lineDigits := make([]int, len(lines))
	for i, line := range lines {
		d, err := GetConcatenatedDigits(line)
		if err != nil {
			return -1, err
		}
		lineDigits[i] = d
	}
	sum := 0
	for _, d := range lineDigits {
		sum += d
	}
	return sum, nil
}
