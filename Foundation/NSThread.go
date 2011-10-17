package foundation

import "github.com/mkrautz/go-objc"

type NSThread struct {
	*objc.Object
}

func NSThreadIsMainThread() bool {
	return objc.GetClass("NSThread").SendMsg("isMainThread") != nil
}
