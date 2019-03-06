package search

import "github.com/gusanmaz/algorithms/types/comparables"

func Linear(a []comparables.Interface, key comparables.Interface) bool{
	for _, v := range a{
		if v.EqualTo(key){
			return true
		}
	}
	return false
}
