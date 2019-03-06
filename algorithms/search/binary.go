package search

import "github.com/gusanmaz/algorithms/types/comparables"

func Binary(a []comparables.Interface, key comparables.Interface) bool{
	lowInd := 0
	highInd := len(a) - 1
	for lowInd <= highInd {
		midInd := lowInd + (highInd - lowInd) / 2
		if a[midInd].EqualTo(key){
			return true
		} else if a[midInd].LessThan(key){
			lowInd = midInd + 1
		} else{
			highInd = midInd - 1
		}
	}
	return false
}
