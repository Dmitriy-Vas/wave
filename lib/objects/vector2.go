package objects

type Vector2 struct {
	X int32
	Y int32
}

func NewVector2(X int32, Y int32) Vector2 {
	return Vector2{
		X: X,
		Y: Y,
	}
}
