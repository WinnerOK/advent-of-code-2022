package utils

type IntVector struct {
	X, Y int
}

func (v IntVector) Add(other IntVector) IntVector {
	return IntVector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v IntVector) Rev() IntVector {
	return IntVector{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v IntVector) Sub(other IntVector) IntVector {
	return v.Add(other.Rev())
}

func (v IntVector) Normalize() IntVector {
	newX := v.X
	if newX != 0 {
		newX = v.X / IntAbs(v.X)
	}
	newY := v.Y
	if newY != 0 {
		newY = v.Y / IntAbs(v.Y)
	}
	return IntVector{
		X: newX,
		Y: newY,
	}
}
