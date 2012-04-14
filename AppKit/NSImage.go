package appkit

import (
	"github.com/mkrautz/objc"
	. "github.com/mkrautz/objc/Foundation"
)

type NSImage struct {
	*objc.Object
}

func NSImageNamed(name string) NSImage {
	return NSImage{objc.GetClass("NSImage").SendMsg("imageNamed:", NSStringFromString(name).Object)}
}
