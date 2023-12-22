package part1

import (
	"slices"
	"testing"
)

const testData = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestParseData(t *testing.T) {
	schematic := ParseData(testData)
	if (len(schematic.lines)) != 10 {
		t.Fail()
	}
	if len(schematic.lines[0]) != 10 {
		t.Fail()
	}
}

func TestGetNeighbors(t *testing.T) {
	schematic := ParseData(testData)
	ns := GetNeighbors(schematic, 0, 4)
	if !slices.Contains(ns, ".") {
		t.Fail()
	}
	if !slices.Contains(ns, "1") {
		t.Fail()
	}
}

func TestGetSchematicSum(t *testing.T) {
	schematic := ParseData(testData)
	s := GetSchematicSum(schematic)
	if s != 4361 {
		t.Fail()
	}
}
