package main

import (
	"github.com/mkrautz/objc"
	. "github.com/mkrautz/objc/AppKit"
	. "github.com/mkrautz/objc/Foundation"
	"runtime"
)

func init() {
	defer runtime.LockOSThread()
}

func main() {
	pool := NewNSAutoreleasePool()
	defer pool.Release()

	// Create a new ObjectiveC class, GOAppDelegate
	c := objc.NewClass(objc.GetClass("NSObject"), "GOAppDelegate")
	// Add a method to it; sayHello
	c.AddMethod("sayHello", objc.EncVoid+objc.EncId+objc.EncSelector)
	// Register the class
	objc.RegisterClass(c)

	// Instantiate the class, and call our new method.
	// For now, this will call an internal method in the objc package.
	o := objc.GetClass("GOAppDelegate").SendMsg("alloc").SendMsg("init")
	o.SendMsg("sayHello")

	app := NSSharedApplication()

	mask := NSTitledWindowMask | NSClosableWindowMask | NSMiniaturizableWindowMask | NSResizableWindowMask
	window := NewNSWindow(NSRectMake(0, 0, 500, 500), mask, NSBackingStoreBuffered, false)
	window.AutoRelease()
	window.Display()
	window.MakeKeyAndOrderFront(window)
	window.SetTitle("Go Demo")

	// Add a gopher!
	gopherImg := NSImageNamed("gopher")
	imgView := NewNSImageView()
	imgView.SetImage(gopherImg)
	window.SetContentView(imgView)

	app.Run()
}
