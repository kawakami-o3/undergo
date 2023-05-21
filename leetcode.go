package u

import (
	"strconv"
	"strings"
	"unicode"
)

func Ints(s string) []int {
	// s = '[1,2,3,...]'
	slice := []int{}
	if len(s) < 2 {
		return slice
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
		slice = append(slice, i)
		digits.Reset()
	}
	return slice
}

func Ints2d(s string) [][]int {
	// s = '[[1,2,3], [4,5,6], ....]'
	slice2d := [][]int{}
	if len(s) < 2 {
		return slice2d
	}

	s = s[1 : len(s)-1]
	slice := []int{}
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
			slice = append(slice, i)
			digits.Reset()
		case '[':
			nest++
			slice = []int{}
		case ']':
			nest--
			if digits.Len() > 0 {
				i, _ := strconv.Atoi(digits.String())
				slice = append(slice, i)
				digits.Reset()
			}
			slice2d = append(slice2d, slice)
		}
	}

	return slice2d
}

func Floats(s string) []float64 {
	// s = '[1,2,3,...]'
	slice := []float64{}
	if len(s) < 2 {
		return slice
	}
	s = s[1:len(s)-1] + string(',')
	var digits strings.Builder
	for _, c := range s {

		switch c {
		case ' ':
			continue
		case ',':
			i, _ := strconv.ParseFloat(digits.String(), 64)
			slice = append(slice, i)
			digits.Reset()
		default:
			digits.WriteRune(c)
		}
	}
	return slice
}

func Floats2d(s string) [][]float64 {
	// s = '[[1,2,3], [4,5,6], ....]'
	slice2d := [][]float64{}
	if len(s) < 2 {
		return slice2d
	}

	s = s[1 : len(s)-1]
	slice := []float64{}
	var digits strings.Builder
	nest := 0
	for _, c := range s {
		switch c {
		case ' ':
			continue
		case ',':
			if nest == 0 {
				continue
			}
			i, _ := strconv.ParseFloat(digits.String(), 64)
			slice = append(slice, i)
			digits.Reset()
		case '[':
			nest++
			slice = []float64{}
		case ']':
			nest--
			if digits.Len() > 0 {
				i, _ := strconv.ParseFloat(digits.String(), 64)
				slice = append(slice, i)
				digits.Reset()
			}
			slice2d = append(slice2d, slice)
		default:
			digits.WriteRune(c)
		}
	}

	return slice2d
}
