package foundation

import (
	"github.com/mkrautz/go-objc"
	"reflect"
	"unsafe"
)

const (
	NSUTF8StringEncoding = 4
)

type NSString struct {
	*objc.Object
}

func NSStringFromString(str string) NSString {
	hdrp := (*reflect.StringHeader)(unsafe.Pointer(&str))
	obj := objc.GetClass("NSString").Alloc().SendMsg("initWithBytes:length:encoding:", hdrp.Data, hdrp.Len, NSUTF8StringEncoding)
	return NSStringFromObject(obj)
}

func NSStringFromObject(obj *objc.Object) NSString {
	return NSString{obj}
}