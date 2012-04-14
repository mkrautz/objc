package appkit

import "github.com/mkrautz/objc"

type NSApplication struct {
	*objc.Object
}

func NSSharedApplication() NSApplication {
	return NSApplication{objc.GetClass("NSApplication").SendMsg("sharedApplication")}
}

func (app NSApplication) Run() {
	app.SendMsg("run")
}
