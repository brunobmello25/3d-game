package noise

import (
	"github.com/aquilax/go-perlin"
)

var p *perlin.Perlin

func Noise2D(x, z int) float64 {
	return getInstance().Noise2D(float64(x)*0.02, float64(z)*0.02) * 20
}

func getInstance() *perlin.Perlin {
	if p == nil {
		p = perlin.NewPerlin(2, 2, 4, 12345)
	}

	return p
}
