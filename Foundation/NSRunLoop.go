package foundation

import "github.com/mkrautz/objc"

type NSRunLoop struct {
	objc.Object
}

func NSRunLoopCurrent() NSRunLoop {
	return NSRunLoop{objc.GetClass("NSRunLoop").SendMsg("currentRunLoop")}
}

func NSRunLoopMain() NSRunLoop {
	return NSRunLoop{objc.GetClass("NSRunLoop").SendMsg("mainRunLoop")}
}

func (rl NSRunLoop) Run() {
	rl.SendMsg("run")
}
