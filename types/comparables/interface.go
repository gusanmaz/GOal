package comparables

type Interface interface {
	LessThan(j Interface) bool // tried (j Comparable), (j MyInt), etc
	EqualTo(j Interface) bool  // tried (j Comparable), (j MyInt), etc
}

