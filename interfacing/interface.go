package interfacing

type Geometry interface {
	Area() float64
	Perim() float64
}

type Rect struct {
	Width, Height float64
}

func (r Rect) Area() float64 {
	return r.Height * r.Width
}

func (r Rect) Perim() float64 {
	return 2*r.Height + 2*r.Width
}

func Measure(g Geometry) (area float64, perim float64) {
	return g.Area(), g.Perim()
}