package part1

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
	fmt.Println("Day 1 Part 1")
	ex, _ := os.Executable()
	f := path.Join(filepath.Dir(ex), "day1", "./input.txt")

	data, err := os.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(SumOfLines(string(data)))
}

func GetFirstAndLastDigits(s string) (string, string) {
	re, _ := regexp.Compile("([0-9])")
	matches := re.FindAllString(s, -1)

	return matches[0], matches[len(matches)-1]
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
