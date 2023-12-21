package part1

import "testing"

func TestParseLine(t *testing.T) {
	id, m := ParseLine("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")
	if id != 3 {
		t.Fail()
	}
	if m[red] != 20 {
		t.Fail()
	}
	if m[green] != 13 {
		t.Fail()
	}
	if m[blue] != 6 {
		t.Fail()
	}
}

func TestIsLineValid(t *testing.T) {
	if !IsLineValid("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green") {
		t.Fail()
	}
	if IsLineValid("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red") {
		t.Fail()
	}
}

func TestSumValidLineIds(t *testing.T) {
	sum := SumValidLineIds(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)
	if sum != 8 {
		t.Fail()
	}
}
