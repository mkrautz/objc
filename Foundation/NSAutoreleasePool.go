package foundation

import "github.com/mkrautz/go-objc"

type NSAutoreleasePool struct {
	*objc.Object
}

func NewNSAutoreleasePool() NSAutoreleasePool {
	return NSAutoreleasePool{objc.GetClass("NSAutoreleasePool").Alloc().Init()}
}
