package comparables

type Float64 float64

func (a Float64) LessThan(b Interface) bool{
	var a1  = float64(a)
	var b1 float64
	switch b.(type){
	case Float64:
		b1 = float64(b.(Float64))
	default:
		b1 = 0
	}
	return a1 < b1
}

func (a Float64) EqualTo(b Interface) bool{
	var a1  = float64(a)
	var b1 float64
	switch b.(type){
	case Float64:
		b1 = float64(b.(Float64))
	default:
		b1 = 0
	}
	return a1 == b1
}

func Float64Arr(a []float64) []Interface {
	retArr := make([]Interface, len(a))
	for i,v := range a{
		retArr[i] = Float64(v)
	}
	return retArr
}

