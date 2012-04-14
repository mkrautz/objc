package appkit

import (
	"github.com/mkrautz/objc"
	. "github.com/mkrautz/objc/Foundation"
)

const (
	NSBorderlessWindowMask         NSUInteger = 0
	NSTitledWindowMask             NSUInteger = 1 << 0
	NSClosableWindowMask           NSUInteger = 1 << 1
	NSMiniaturizableWindowMask     NSUInteger = 1 << 2
	NSResizableWindowMask          NSUInteger = 1 << 3
	NSTexturedBackgroundWindowMask NSUInteger = 1 << 8
)

const (
	NSBackingStoreRetained    = 0
	NSBackingStoreNonretained = 1
	NSBackingStoreBuffered    = 2
)

type NSBackingStoreType NSUInteger

type NSWindow struct {
	*objc.Object
}

func NewNSWindow(rect NSRect, windowStyle NSUInteger, bufferingType NSBackingStoreType, deferCreation bool) NSWindow {
	obj := objc.GetClass("NSWindow").Alloc().
		SendMsg("initWithContentRect:styleMask:backing:defer:",
		rect, windowStyle, bufferingType, deferCreation)
	return NSWindow{obj}
}

func (win NSWindow) Display() {
	win.SendMsg("display")
}

func (win NSWindow) MakeKeyAndOrderFront(sender *objc.Object) {
	win.SendMsg("makeKeyAndOrderFront:", sender)
}

func (win NSWindow) SetTitle(title string) {
	win.SendMsg("setTitle:", NSStringFromString(title).Object)
}

func (win NSWindow) Title() string {
	return win.SendMsg("title").String()
}