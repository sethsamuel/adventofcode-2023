package part1

import "testing"

func TestGetFirstAndLastDigits(t *testing.T) {
	first, last := GetFirstAndLastDigits("as2lkj3lk8lkj")
	if first != "2" || last != "8" {
		t.Fail()
	}
}

func TestGetConcatenatedDigits(t *testing.T) {
	i, err := GetConcatenatedDigits("as2lkj3lk8lkj")
	if err != nil {
		t.Fail()
	}
	if i != 28 {
		t.Fail()
	}
}

func TestSumOfLines(t *testing.T) {
	i, err := SumOfLines(`1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`)

	if err != nil {
		t.Fail()
	}
	if i != 142 {
		t.Fail()
	}
}
