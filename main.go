package main

import (
	"fmt"
	"hack20-35-fuzzing/strfuncs"
	"hack20-35-fuzzing/interfacing"
)

// based on https://medium.com/a-journey-with-go/go-fuzz-testing-in-go-deb36abc971f

func main() {
	r := strfuncs.CountChars("Adam")
	fmt.Printf("%v characters\n", r)

	newShape := interfacing.Rect{
		Width:  2.44,
		Height: 2.44545,
	}

	a, p := interfacing.Measure(newShape)
	fmt.Printf("Area: %f, Perimeter: %f\n", a, p)
}
