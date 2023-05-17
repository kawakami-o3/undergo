package u

import (
	"strconv"
	"strings"
	"unicode"
)

func Ints(s string) []int {
	// s = '[1,2,3,...]'
	ints := []int{}
	if len(s) < 2 {
		return ints
	}
	delimiter := ','
	s = s[1:len(s)-1] + string(delimiter)
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
	// s = '[[1,2,3], [4,5,6], ....]'
	ints2d := [][]int{}
	maxNext := maxNext(s)
	if maxNext == 1 || maxNext == 2 {
		return ints2d // invalid structure
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
