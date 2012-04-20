package appkit

import "github.com/mkrautz/objc"

type NSApplication struct {
	objc.Object
}

func NSSharedApplication() NSApplication {
	return NSApplication{objc.GetClass("NSApplication").SendMsg("sharedApplication")}
}

func (app NSApplication) Run() {
	app.SendMsg("run")
}

func (app NSApplication) SetDelegate(delegate objc.Object) {
	app.SendMsg("setDelegate:", delegate)
}

func (app NSApplication) Delegate() objc.Object {
	return app.SendMsg("delegate")
}
