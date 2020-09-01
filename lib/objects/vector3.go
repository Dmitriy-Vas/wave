package objects

type Vector3 struct {
	X, Y, Z int32
}

func NewVector3(x, y, z int32) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}
