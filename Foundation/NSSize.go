package foundation

import "fmt"

type NSSize struct {
	Width  float32
	Height float32
}

func (sz NSSize) String() string {
	return fmt.Sprintf("(%v, %v)", sz.Width, sz.Height)
}
