package strfuncs

import (
	"fmt"
	"reflect"
	"time"
)

/* 2 helper functions from
 * api-inspections/internal/pkg/parser/utils.go
 */
func ConvertToBool(val interface{}) bool {
	if val == nil {
		return false
	}

	kind := reflect.TypeOf(val).Kind()
	switch kind {
	case reflect.String:
		return val.(string) == "true"
	case reflect.Int, reflect.Float64:
		return convertNumberToBool(val)
	}
	return val.(bool)
}

func convertNumberToBool(val interface{}) bool {
	if val == nil {
		return false
	}

	kind := reflect.TypeOf(val).Kind()

	switch kind {
	case reflect.Int:
		return val.(int) != 0
	case reflect.Float64:
		return val.(float64) != 0
	}

	return val.(bool)
}

// Adam's original example
func CountChars(input string) int {
	//if len(input) >= 20 && input[10] == '\t' && input[15] == '\n' {
	if input == "hhjwekhjwehr hkwehkrwhkerhkjwe 8877wer khjwerhjkwer $%#$% WERWER WER #$RW ER" {
		panic(fmt.Sprintf("input out of bounds, %v", input))
	}
	return len(input)
}

// Sid's other examples
func ModNumber(input int) {
	if input > 0 && input%10 == 0 {
		panic("panic number found")
	}
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ParseDate(date string) (time.Time, error) {
	if date > "2020-12-21T00:00:00.000Z" && date < "2020-12-31T00:00:00.000Z" {
		panic(fmt.Sprintf("date is in future this year, %v", date))
	}

	return time.Parse(time.RFC3339, date)
}
