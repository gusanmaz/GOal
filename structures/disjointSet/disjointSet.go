package disjointSet

type Algorithm int

const (
	FastFind = iota
	FastUnion
	WeightUnion
	PathCompression
)

type DisJointSet struct {
	algorithm Algorithm
	size      int
	cnt       int
	id        []int
	sz        []int
}

type Props struct {
	Algorithm Algorithm
	Size      int
}

func New(props Props) *DisJointSet {
	if props.Size == 0 {
		// Default Size to 10
		props.Size = 10
	}

	switch props.Algorithm {
		case FastFind, FastUnion, WeightUnion, PathCompression:
		default:
			props.Algorithm = FastFind
	}

	id := make([]int, props.Size)
	for i, _ := range id{
		id[i] = i
	}

	var sz []int
	if props.Algorithm == WeightUnion || props.Algorithm == PathCompression {
		sz = make([]int, props.Size)
	}

	return &DisJointSet{props.Algorithm, props.Size, props.Size, id, sz}
}

func (s *DisJointSet) isIndexValid(ind int) bool{
	if ind < 0 || ind >= s.size{
		return false
	}
	return true
}

func (s *DisJointSet) Connected(p, q int) bool {
	pId := s.Find(p)
	qId := s.Find(q)
	if pId == qId{
		return true
	}
	return false
}

func (s *DisJointSet) Union(p, q int) {
	if !s.isIndexValid(p) || !s.isIndexValid(q){
		// may panic()
		return
	}
	if s.algorithm == FastFind{
		pId := s.Find(p)
		qId := s.Find(q)
		if pId != qId{
			for  i,_ := range s.id{
				if s.id[i] == qId{
					s.id[i] = pId
					s.cnt--
				}
			}
		}
	}
}

func (s *DisJointSet) Find(elem int) int {
	if !s.isIndexValid(elem){
		// may panic
		return -1
	}
	if s.algorithm == FastFind{
		return s.id[elem]
	}
	// TODO
	return -1
}
