package objc

import (
	"sync"
	"testing"
)

type SomeObject struct {
	Object

	t *testing.T
}

const (
	float64val float64 = 42.0
)

var once sync.Once

func registerTestClass() {
	once.Do(func() {
		c := NewClass(GetClass("NSObject"), "SomeObject", SomeObject{})
		c.AddMethod("callWithObject:selector:", (*SomeObject).CallWithObjectAndSelector)
		RegisterClass(c)
	})
}

func (so *SomeObject) CallWithObjectAndSelector(object Object, selector Selector) {
	if selector.Selector() != "callWithObject:selector:" {
		so.t.Errorf("unexpected selector")
	}
	if so.Pointer() != object.Pointer() {
		so.t.Errorf("unexpected object")
	}
}

func TestSelectorObjectPassing(t *testing.T) {
	registerTestClass()
	so := &SomeObject{t: t}
	NewGoInstance("SomeObject", so)
	so.SendMsg("callWithObject:selector:", so, GetSelector("callWithObject:selector:"))
}