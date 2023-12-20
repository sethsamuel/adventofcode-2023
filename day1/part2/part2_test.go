package part2

import "testing"

func TestGetFirstAndLastDigits(t *testing.T) {
	first, last := GetFirstAndLastDigits("aone2sevenx")
	if first != "1" || last != "7" {
		t.Fail()
	}
}

func TestGetConcatenatedDigits(t *testing.T) {
	i, err := GetConcatenatedDigits("as2lkj3lk8lninekj")
	if err != nil {
		t.Fail()
	}
	if i != 29 {
		t.Fail()
	}
}

func TestSumOfLines(t *testing.T) {
	i, err := SumOfLines(`two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`)

	if err != nil {
		t.Fail()
	}
	if i != 281 {
		t.Fail()
	}
}
