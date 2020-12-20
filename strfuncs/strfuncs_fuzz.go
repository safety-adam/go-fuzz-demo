package strfuncs

func FuzzCountChars(data []byte) int {
	s := string(data)
	r := CountChars(s)

	if len(s) == r {
		return 1
	}

	return -1
}
