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

func Ints2d(s string) [][]int {
	// s = '[[1,2,3], [4,5,6], ....]'
	ints2d := [][]int{}
	if len(s) < 2 {
		return ints2d
	}

	s = s[1 : len(s)-1]
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
			if nest == 0 {
				continue
			}
			i, _ := strconv.Atoi(digits.String())
			ints = append(ints, i)
			digits.Reset()
		case '[':
			nest++
			ints = []int{}
		case ']':
			nest--
			if digits.Len() > 0 {
				i, _ := strconv.Atoi(digits.String())
				ints = append(ints, i)
				digits.Reset()
			}
			ints2d = append(ints2d, ints)
		}
	}

	return ints2d
}
