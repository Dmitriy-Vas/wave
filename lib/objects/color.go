package objects

type Color struct {
	Red, Green, Blue, Alpha byte
}

func NewColor(red, green, blue, alpha byte) Color {
	return Color{
		Red:   red,
		Green: green,
		Blue:  blue,
		Alpha: alpha,
	}
}
