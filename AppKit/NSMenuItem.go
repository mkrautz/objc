package appkit

import (
	"github.com/mkrautz/objc"
	. "github.com/mkrautz/objc/Foundation"
)

type NSMenuItem struct {
	objc.Object
}

func NewNSMenuItem(itemName string, action objc.Selector, keyEquivalent string) NSMenuItem {
	return NSMenuItem{objc.GetClass("NSMenuItem").SendMsg("alloc").SendMsg("initWithTitle:action:keyEquivalent:",
		NSStringFromString(itemName), action, NSStringFromString(keyEquivalent))}
}