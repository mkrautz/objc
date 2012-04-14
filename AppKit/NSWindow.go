package appkit

import (
	"github.com/mkrautz/objc"
	. "github.com/mkrautz/objc/Foundation"
)

const (
	NSBorderlessWindowMask         = 0
	NSTitledWindowMask             = 1 << 0
	NSClosableWindowMask           = 1 << 1
	NSMiniaturizableWindowMask     = 1 << 2
	NSResizableWindowMask          = 1 << 3
	NSTexturedBackgroundWindowMask = 1 << 8
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
