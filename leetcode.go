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
		if unicode.IsDigit(c) || c == '-' {
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
		if unicode.IsDigit(c) || c == '-' {
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

func Strings(s string) []string {
	slice := []string{}
	if len(s) < 2 {
		return slice
	}
	s = s[1:len(s)-1] + string(',')
	var buf strings.Builder
	for _, c := range s {

		switch c {
		case ' ', '"':
			continue
		case ',':
			slice = append(slice, buf.String())
			buf.Reset()
		default:
			buf.WriteRune(c)
		}
	}
	return slice
}

func Strings2d(s string) [][]string {
	slice2d := [][]string{}
	if len(s) < 2 {
		return slice2d
	}

	s = s[1 : len(s)-1]
	slice := []string{}
	var buf strings.Builder
	nest := 0
	for _, c := range s {
		switch c {
		case ' ', '"':
			continue
		case ',':
			if nest == 0 {
				continue
			}
			slice = append(slice, buf.String())
			buf.Reset()
		case '[':
			nest++
			slice = []string{}
		case ']':
			nest--
			if buf.Len() > 0 {
				slice = append(slice, buf.String())
				buf.Reset()
			}
			slice2d = append(slice2d, slice)
		default:
			buf.WriteRune(c)
		}
	}

	return slice2d
}

func Chars(s string) []byte {
	// s = `["a","b","c",...]`
	slice := []byte{}
	if len(s) < 2 {
		return slice
	}
	reader := strings.NewReader(s)
	for reader.Len() > 0 {
		b, _ := reader.ReadByte()
		if b == '\'' || b == '"' {
			c, _ := reader.ReadByte()
			slice = append(slice, byte(c))
			reader.ReadByte()
		}
	}
	return slice
}
