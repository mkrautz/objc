package objc

import "testing"

type ExampleStruct struct {
	a uint32
	b uint32
	c uint64
}

func (e ExampleStruct) PackStruct64() (args []uint64) {
	args = make([]uint64, 2)	
	args[0] = uint64(e.a) << 32 | uint64(e.b) 
	args[1] = uint64(e.c)
	return args
}

func (e ExampleStruct) PackStruct32() (args []uint32) {
	panic("unimplemented")
}

func TestExampleStruct(t *testing.T) {
	e := ExampleStruct{
		a: 0x11111111,
		b: 0x22222222,
		c: 0x3333333333333333,
	}
	args := e.PackStruct64()
	if args[0] != 0x1111111122222222 {
		t.Error("args[0] packed wrong")
	}
	if args[1] != 0x3333333333333333 {
		t.Error("args[1] packed wrong")
	}
}
