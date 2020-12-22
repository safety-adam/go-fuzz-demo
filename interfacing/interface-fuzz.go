// +build gofuzz
package interfacing
import(
  "encoding/json"
)

func FuzzInterfacing(data []byte) int {
	r := &Rect{}

  err := json.Unmarshal(data, r)
  if err != nil {
    return -1
  }

  expected := r.Width * r.Height

	area := r.Area()
  if (area != expected) {
    panic("Area: is not equal and has caused problem!")
  }
	return 1
}
