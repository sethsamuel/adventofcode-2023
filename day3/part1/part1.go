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

func Run() {

	schematic := ParseData(data)
	fmt.Println(GetSchematicSum(schematic))

}

type schematic struct {
	lines [][]string
}

func ParseData(data string) schematic {
	lines := strings.Split(data, "\n")
	splitLines := make([][]string, len(lines))
	for i, l := range lines {
		splitLines[i] = strings.Split(l, "")
	}
	return schematic{lines: splitLines}
}

func GetValue(schematic schematic, x int, y int) string {
	return schematic.lines[y][x]
}

func GetOffset(schematic schematic, x int, y int, xOffset int, yOffset int) string {
	if x+xOffset < 0 {
		return ""
	}
	if y+yOffset < 0 {
		return ""
	}
	if x+xOffset >= len(schematic.lines[0]) {
		return ""
	}
	if y+yOffset >= len(schematic.lines) {
		return ""
	}
	return GetValue(schematic, x+xOffset, y+yOffset)
}

func GetNeighbors(schematic schematic, x int, y int) []string {
	ns := make([]string, 8)
	n := 0

	for xOffset := -1; xOffset <= 1; xOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			if xOffset == 0 && yOffset == 0 {
				continue
			}
			ns[n] = GetOffset(schematic, x, y, xOffset, yOffset)
			n++
		}
	}
	return ns
}

type OptionalBoolean int

const (
	Unknown OptionalBoolean = iota
	True
	False
)

func GetSchematicSum(schematic schematic) int {
	s := 0
	re := regexp.MustCompile("[0-9]")
	for y, l := range schematic.lines {
		currentNumber := ""
		isAdjacentToSymbol := Unknown
		for x, _ := range l {
			v := GetValue(schematic, x, y)

			isNumber := re.MatchString(v)
			if !isNumber {
				if isAdjacentToSymbol == True {

					i, _ := strconv.Atoi(currentNumber)
					s += i
				}
				isAdjacentToSymbol = Unknown
				currentNumber = ""
				continue
			}
			currentNumber += v
			if isAdjacentToSymbol == Unknown {
				ns := GetNeighbors(schematic, x, y)

				for _, n := range ns {
					if n != "." && n != "" && !re.MatchString(n) {
						isAdjacentToSymbol = True
						break
					}
				}
			}

		}

		if isAdjacentToSymbol == True {

			i, _ := strconv.Atoi(currentNumber)
			s += i
		}
	}

	return s
}
