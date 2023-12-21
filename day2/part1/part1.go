package part1

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed "input.txt"
var data string

type Color string

const (
	red   Color = "red"
	green Color = "green"
	blue  Color = "blue"
)

var limits = map[Color]int{
	red:   12,
	green: 13,
	blue:  14,
}

func Run() {
	fmt.Println("Day 2 Part 1")
	fmt.Println(SumValidLineIds(data))
}

// Example line:
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
func ParseLine(s string) (int, map[Color]int) {
	re, _ := regexp.Compile("Game ([0-9]+):")
	matches := re.FindStringSubmatch(s)
	id, _ := strconv.Atoi(matches[1])

	m := make(map[Color]int)
	draws := regexp.MustCompile("[;,]").Split(strings.TrimSpace(strings.Split(s, ":")[1]), -1)
	for _, s := range draws {
		parts := strings.Split(strings.TrimSpace(s), " ")
		count, _ := strconv.Atoi(parts[0])
		color := Color(parts[1])
		m[color] = max(count, m[color])
	}

	return id, m
}

func IsLineValid(s string) bool {
	_, counts := ParseLine(s)
	return counts[red] <= limits[red] && counts[green] <= limits[green] && counts[blue] <= limits[blue]
}

func SumValidLineIds(s string) int {
	lines := strings.Split(s, "\n")
	validIdSum := 0
	for _, s := range lines {
		id, _ := ParseLine(s)
		if IsLineValid(s) {
			validIdSum += id
		}
	}
	return validIdSum
}

// func GetConcatenatedDigits(s string) (int, error) {
// 	first, last := GetFirstAndLastDigits(s)
// 	return strconv.Atoi(first + last)
// }

// func SumOfLines(s string) (int, error) {
// 	lines := strings.Split(s, "\n")
// 	lineDigits := make([]int, len(lines))
// 	for i, line := range lines {
// 		d, err := GetConcatenatedDigits(line)
// 		if err != nil {
// 			return -1, err
// 		}
// 		lineDigits[i] = d
// 	}
// 	sum := 0
// 	for _, d := range lineDigits {
// 		sum += d
// 	}
// 	return sum, nil
// }