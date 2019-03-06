package generators

import (
	"math/rand"
	"time"
)

type SFloat64ArrProps struct {
	Size uint32
}

func SFloat64Arr(p *SFloat64ArrProps) []float64 {
	var size uint32 = 20
	if p.Size != 0{
		size = p.Size
	}

	rand.Seed(time.Now().Unix())
	randArr := make([]float64, size)

	for i, _ :=  range randArr{
		r := rand.Float64()
		negate := rand.Uint32() % 2
		if negate == 1{
			r *= -1
		}

		randArr[i] = r
	}
	return randArr
}

