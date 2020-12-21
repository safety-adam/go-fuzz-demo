// +build gofuzz
package interfacing

import fuzz "github.com/google/gofuzz"

func FuzzInterfacing(data []byte) int {
	var shape Rect
	fuzz.NewFromGoFuzz(data).Fuzz(&shape)
	shape.Area()
	return 0
}
