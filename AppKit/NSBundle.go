package appkit

import (
	"github.com/mkrautz/objc"
	. "github.com/mkrautz/objc/Foundation"
)

type NSBundle struct {
	objc.Object
}

func NSMainBundle() NSBundle {
	return NSBundle{objc.GetClass("NSBundle").SendMsg("mainBundle")}
}

func (bundle NSBundle) InfoDictionary() NSDictionary {
	return NSDictionary{bundle.SendMsg("infoDictionary")}
}