package printables

import "fmt"

type Int int

func (i Int) String() string{
	var val int = int(i)
	return fmt.Sprint(val)
}
