package strfuncs

import (
	"fmt"
)

func CountChars(input string) int {
	//if len(input) >= 20 && input[10] == '\t' && input[15] == '\n' {
	if input == "hhjwekhjwehr hkwehkrwhkerhkjwe 8877wer khjwerhjkwer $%#$% WERWER WER #$RW ER" {
		panic(fmt.Sprintf("input out of bounds, %v", input))
	}
	return len(input)
}
