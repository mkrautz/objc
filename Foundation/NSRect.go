package foundation

import "fmt"

type NSRect struct {
	Origin NSPoint
	Size   NSSize
}

func (r NSRect) PackStruct64() (args []uint64) {
	args = make([]uint64, 2)
	args[0] = uint64(r.Origin.X)<<32 | uint64(r.Origin.Y)
	args[1] = uint64(r.Size.Width)<<32 | uint64(r.Size.Height)
	return args
}

func (r NSRect) PackStruct32() (args []uint32) {
	panic("unimplemented")
}

func (r NSRect) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", r.Origin.X, r.Origin.Y, r.Size.Width, r.Size.Height)
}

func NSRectMake(x, y, w, h float32) NSRect {
	return NSRect{
		NSPoint{
			x, y,
		},
		NSSize{
			w, h,
		},
	}
}
