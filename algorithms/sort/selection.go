package sort

import (
	"github.com/gusanmaz/algorithms/helper"
	"github.com/gusanmaz/algorithms/types/comparables"
)

func Selection(a []comparables.Interface){
	for i,_ := range a{
		minInd := i
		for k := i + 1; k < len(a); k++{
			if helper.Comp(a, minInd, k) > 0{
				minInd = k
			}
		}
		helper.Swap(a, minInd, i)
	}
}
