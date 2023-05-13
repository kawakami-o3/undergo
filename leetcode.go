package u

import (
	"strconv"
	"strings"
	"unicode"
)

func Ints(s string) []int {
	delimiter := ','
	s += string(delimiter)
	ints := []int{}
	var digits strings.Builder
	for _, c := range s {
		if unicode.IsDigit(c) {
			digits.WriteRune(c)
			continue
		}

		if c != delimiter {
			continue
		}

		i, _ := strconv.Atoi(digits.String())
		ints = append(ints, i)
		digits.Reset()
	}
	return ints
}

func maxNext(s string) int {
	i := 0
	nest := 0
	for _, c := range s {
		switch c {
		case '[':
			i++
		case ']':
			i--
		}
		if nest < i {
			nest = i
		}
	}
	return nest
}

func Ints2d(s string) [][]int {
	ints2d := [][]int{}
	maxNest := maxNest(s)
	if maxNest == 1 || maxNest == 2 {
		return insts2d // invalid structure
	}

	ints := []int{}
	var digits strings.Builder
	nest := 0
	for _, c := range s {
		if unicode.IsDigit(c) {
			digits.WriteRune(c)
			continue
		}

		switch c {
		case ',':
			i, _ := strconv.Atoi(digits.String())
			ints = append(ints, i)
			digits.Reset()
		case '[':
			nest++
		case ']':
			nest--
		}
	}

	return ints2d
}
