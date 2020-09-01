package objects

type Vector4 struct {
	A, B, C, D int32
}

func NewVector4(a, b, c, d int32) Vector4 {
	return Vector4{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}
