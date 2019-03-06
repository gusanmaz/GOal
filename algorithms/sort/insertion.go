package sort

import (
	"github.com/gusanmaz/algorithms/helper"
	"github.com/gusanmaz/algorithms/types/comparables"
)


func insertion(a []comparables.Interface){
	for i := 1; i < len(a); i++{
		for k := i; k > 0; k--{
			if helper.Comp(a, k, k - 1) < 0 {
				helper.Swap(a, k, k -1)
			}
		}
	}
}
