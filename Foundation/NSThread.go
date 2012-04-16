package foundation

import "github.com/mkrautz/objc"

type NSThread struct {
	objc.Object
}

func NSThreadIsMainThread() bool {
	return objc.GetClass("NSThread").SendMsg("isMainThread") != nil
}
