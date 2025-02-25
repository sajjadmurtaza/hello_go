package math

type Rectangle struct {
	Width  float32
	Length float32
}

func (r Rectangle) GetArea() float32 {
	return r.Width * r.Length
}
