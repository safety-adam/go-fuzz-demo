package strfuncs

import (
	"strconv"
	"strings"
)

/* fuzz test for 2 helper functions from
 * api-inspections/internal/pkg/parser/utils.go
 */
func FuzzConvertToBool(data []byte) int {
	s := string(data)
	if ConvertToBool(s) == false {
		if strings.ToLower(s) == "true" {
			panic("true despite returning false")
		}
	}

	return 1
}

// Adam's original example
func FuzzCountChars(data []byte) int {
	s := string(data)
	r := CountChars(s)

	if len(s) == r {
		return 1
	}

	return -1
}

// Sid's other examples
func FuzzModNumber(data []byte) int {
	s := string(data)
	i, err := strconv.Atoi(s)
	if err == nil {
		ModNumber(i)
		return 1
	}

	return -1
}

func FuzzReverseString(data []byte) int {
	s := string(data)
	if s == ReverseString(ReverseString(s)) {
		return 1
	}

	panic("reverse didn't work")
}

func FuzzParseDate(data []byte) int {
	s := string(data)
	time, error := ParseDate(s)
	if error != nil {
		if !time.IsZero() {
			panic("time != zero on error")
		}
	}

	return 1
}
