package objc

/*
#cgo LDFLAGS: -lobjc
#include <objc/runtime.h>

void *GoObjc_RegisterSelector(char *name) {
	return (void *) sel_registerName(name);
}
*/
import "C"
import "unsafe"

// selectorWithName looks up a selector by name.
func selectorWithName(name string) unsafe.Pointer {
	return C.GoObjc_RegisterSelector(C.CString(name))
}
