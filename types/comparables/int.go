package comparables

type Int int

func (a Int) LessThan(b Interface) bool{
	var a1  = int(a)
	var b1 int
	switch b.(type){
	case Int:
		b1 = int(b.(Int))
	default:
		b1 = 0
	}
	return a1 < b1
}

func (a Int) EqualTo(b Interface) bool{
	var a1  = int(a)
	var b1 int
	switch b.(type){
	case Int:
		b1 = int(b.(Int))
	default:
		b1 = 0
	}
	return a1 == b1
}

func IntArr(a []int) []Interface {
	retArr := make([]Interface, len(a))
	for i,v := range a{
		retArr[i] = Int(v)
	}
	return retArr
}
