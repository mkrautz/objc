package main

import (
	"github.com/mkrautz/objc"
	. "github.com/mkrautz/objc/AppKit"
	. "github.com/mkrautz/objc/Foundation"
	"log"
	"runtime"
)

func init() {
	defer runtime.LockOSThread()
}

type AppDelegate struct {
	objc.Object
}

func init() {
	c := objc.NewClass(objc.GetClass("NSObject"), "GOAppDelegate", AppDelegate{})
	c.AddMethod("applicationDidFinishLaunching:", (*AppDelegate).ApplicationDidFinishLaunching)
	objc.RegisterClass(c)
}

func NewAppDelegate() *AppDelegate {
	appDelegate := new(AppDelegate)
	objc.NewGoInstance("GOAppDelegate", appDelegate)
	return appDelegate
}

func (delegate *AppDelegate) ApplicationDidFinishLaunching(notification objc.Object) {
	log.Printf("applicationDidFinishLaunching! %v", notification)
}

func main() {
	pool := NewNSAutoreleasePool()
	defer pool.Release()

	appDelegate := NewAppDelegate()

	app := NSSharedApplication()
	app.SetDelegate(appDelegate)

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
