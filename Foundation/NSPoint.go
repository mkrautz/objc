package foundation

import "fmt"

type NSPoint struct {
	X float32
	Y float32
}

func (pt NSPoint) String() string {
	return fmt.Sprintf("(%v, %v)", pt.X, pt.Y)
}
