package utils

type Dimensions struct {
	X int
	Y int
}

func NewDimensions(x, y int) Dimensions {
	return Dimensions{
		X: x,
		Y: y,
	}
}
