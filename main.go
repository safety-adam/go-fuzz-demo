package main

import (
	"fmt"
	"hack20-35-fuzzing/strfuncs"
)

// based on https://medium.com/a-journey-with-go/go-fuzz-testing-in-go-deb36abc971f

func main() {
	r := strfuncs.CountChars("Adam")
	fmt.Printf("%v characters\n", r)
}
