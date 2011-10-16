package objc

import (
	"log"
	"reflect"
	"testing"
	"unsafe"
)

func TestGetClass(t *testing.T) {
	helloWorld := "hello world from Go!"
	hdrp := (*reflect.StringHeader)(unsafe.Pointer(&helloWorld))
	obj := GetClass("NSString").Alloc().SendMsg("initWithBytes:length:encoding:", hdrp.Data, hdrp.Len, NSUTF8StringEncoding)
	log.Printf("%v", obj)
}