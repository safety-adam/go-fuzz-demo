// +build gofuzz
package strfuncs

import fuzz "github.com/google/gofuzz"

func FuzzCountChars(data []byte) int {
	var s string
	fuzz.NewFromGoFuzz(data).Fuzz(&s)
	CountChars(s)
	return 0
}
