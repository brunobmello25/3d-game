package noise

import "math"

func Noise2D(x, z int) float64 {
	return (math.Sin(float64(x)*0.1) + math.Cos(float64(z)*0.1)) * 10
}
