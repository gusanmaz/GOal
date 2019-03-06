package printables

import "fmt"

type Intp int

func (i *Intp) String() string{
	var val int = int(*i)
	return fmt.Sprint(val)
}
