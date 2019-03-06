package generators

import (
	"math"
	"math/rand"
	"time"
)

type IntArrProps struct {
	Min  int32
	Max  int32 // Exclusive
	Size uint32
}

func IntArr(p *IntArrProps) []int32 {
	var min int32 = math.MinInt32
	var max int32 = math.MaxInt32
	var size uint32 = 20

	if p.Min != 0 {
		min = p.Min
	}

	if p.Max != 0 {
		max = p.Max
	}

	if p.Size != 0{
		size = p.Size
	}

	rand.Seed(time.Now().Unix())
	randArr := make([]int32, size)
	diff := uint32(max - min)

	for i, _ :=  range randArr{
		r := rand.Uint32()
		add := r % diff
		randArr[i] = min + int32(add)

	}
	return randArr
}
