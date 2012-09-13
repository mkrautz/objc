// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/mkrautz/objc"
	. "github.com/mkrautz/objc/AppKit"
	"log"
	"runtime"
)

func init() {
	defer runtime.LockOSThread()

	c := objc.NewClass(objc.GetClass("NSObject"), "GOAppDelegate", AppDelegate{})
	c.AddMethod("applicationDidFinishLaunching:", (*AppDelegate).ApplicationDidFinishLaunching)
	objc.RegisterClass(c)
}

type AppDelegate struct {
	objc.Object
}

func NewAppDelegate() *AppDelegate {
	appDelegate := new(AppDelegate)
	objc.NewGoInstance("GOAppDelegate", appDelegate)
	return appDelegate
}

func (delegate *AppDelegate) ApplicationDidFinishLaunching(notification objc.Object) {
	log.Printf("applicationDidFinishLaunching! %v", notification)

	mainMenu := NSSharedApplication().MainMenu()
	log.Printf("%v", mainMenu)
}

func main() {
	NSApplicationMain()
	/*pool := NewNSAutoreleasePool()
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

	app.Run()*/
}
