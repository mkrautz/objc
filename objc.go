// Package objc implements access to the Objective-C runtime from Go
package objc

/*
#cgo LDFLAGS: -lobjc -framework Foundation
#include <objc/runtime.h>
#include <objc/message.h>

#define ulong unsigned long

void *GoObjc_GetClass(char *name) {
	return (void *) objc_getClass(name);
}

void *GoObjc_SendMsg0(void *receiver, void *selector) {
	return (void *) objc_msgSend(receiver, selector);
}

void *GoObjc_SendMsg1(void *receiver, void *selector, ulong a1) {
	return (void *) objc_msgSend(receiver, selector, a1);
}

void *GoObjc_SendMsg2(void *receiver, void *selector, ulong a1, ulong a2) {
	return (void *) objc_msgSend(receiver, selector, a1, a2);
}

void *GoObjc_SendMsg3(void *receiver, void *selector, ulong a1, ulong a2, ulong a3) {
	return (void *) objc_msgSend(receiver, selector, a1, a2, a3);
}

void *GoObjc_SendMsg4(void *receiver, void *selector, ulong a1, ulong a2, ulong a3, ulong a4) {
	return (void *) objc_msgSend(receiver, selector, a1, a2, a3, a4);
}

void *GoObjc_SendMsg5(void *receiver, void *selector, ulong a1, ulong a2, ulong a3, ulong a4, ulong a5) {
	return (void *) objc_msgSend(receiver, selector, a1, a2, a3, a4, a5);
}

void *GoObjc_SendMsg6(void *receiver, void *selector, ulong a1, ulong a2, ulong a3, ulong a4, ulong a5, ulong a6) {
	return (void *) objc_msgSend(receiver, selector, a1, a2, a3, a4, a5, a6);
}

void *GoObjc_RegisterSelector(char *name) {
	return (void *) sel_registerName(name);
}
*/
import "C"
import (
	"unsafe"
)

// A Selector represents an Objective-C selector.
type Selector uintptr

// Look up a selector by its name
func SelectorName(name string) Selector {
	return Selector(C.GoObjc_RegisterSelector(C.CString(name)))
}

// Checks whether the Selector s is nil.
func (s Selector) IsNil() bool {
	return uintptr(s) == 0 
}

// A Class represents an Objective-C class.
type Class struct {
}

// An Object represents an Objective-C object.
type Object struct {
	isa *Class
}

// Lookup a Class by name
func GetClass(name string) *Object {
	 return (*Object)(C.GoObjc_GetClass(C.CString(name)))
}

// Return the Object as a uintptr.
// Using package unsafe, this uintptr can further be converted to an unsafe.Pointer.
func (obj *Object) Pointer() uintptr {
	return uintptr(unsafe.Pointer(obj))
}

// Send a message to an object.
func (obj *Object) SendMsg(selectorName string, args ...interface{}) *Object {
	// Keep ObjC semantics: messages can be sent to nil objects,
	// but the response is nil.
	if obj == nil {
		return nil
	}

	sel := SelectorName(selectorName)
	if sel.IsNil() {
		return nil
	}

	passableArgs := make([]uintptr, len(args))
	for i, arg := range args {
		switch t := arg.(type) {
		case uintptr:
			passableArgs[i] = t
			break
		case int:
			passableArgs[i] = uintptr(t)
		default:
			panic("unhandled kind")
		}	
	}

	switch len(args) {
	case 0:
		return (*Object)(C.GoObjc_SendMsg0(unsafe.Pointer(obj), unsafe.Pointer(sel)))
	case 1:
		return (*Object)(C.GoObjc_SendMsg1(unsafe.Pointer(obj), unsafe.Pointer(sel),
							C.ulong(passableArgs[0])))
	case 2:
		return (*Object)(C.GoObjc_SendMsg2(unsafe.Pointer(obj), unsafe.Pointer(sel),
							C.ulong(passableArgs[0]), C.ulong(passableArgs[1])))
	case 3:
		return (*Object)(C.GoObjc_SendMsg3(unsafe.Pointer(obj), unsafe.Pointer(sel),
							C.ulong(passableArgs[0]), C.ulong(passableArgs[1]),
							C.ulong(passableArgs[2])))
	case 4:
		return (*Object)(C.GoObjc_SendMsg4(unsafe.Pointer(obj), unsafe.Pointer(sel),
							C.ulong(passableArgs[0]), C.ulong(passableArgs[1]),
							C.ulong(passableArgs[2]), C.ulong(passableArgs[3])))
	case 5:
		return (*Object)(C.GoObjc_SendMsg5(unsafe.Pointer(obj), unsafe.Pointer(sel),
							C.ulong(passableArgs[0]), C.ulong(passableArgs[1]),
							C.ulong(passableArgs[2]), C.ulong(passableArgs[3]),
							C.ulong(passableArgs[4])))
	case 6:
		return (*Object)(C.GoObjc_SendMsg6(unsafe.Pointer(obj), unsafe.Pointer(sel),
							C.ulong(passableArgs[0]), C.ulong(passableArgs[1]),
							C.ulong(passableArgs[2]), C.ulong(passableArgs[3]),
							C.ulong(passableArgs[4]), C.ulong(passableArgs[5])))
	}

	panic("unimplemented amount of SendMsg args")
}

// Send the "alloc" message to the Object.
func (obj *Object) Alloc() *Object {
	return obj.SendMsg("alloc")
}

// Send the "init" message to the Object.
func (obj *Object) Init() *Object {
	return obj.SendMsg("init")
}

// Send the "retain" message to the Object.
func (obj *Object) Retain() *Object {
	return obj.SendMsg("retain")
}

// Send the "release" message to the Object.
func (obj *Object) Release() *Object {
	return obj.SendMsg("release")
}

// Send the "autorelease" message to the Object.
func (obj *Object) AutoRelease() *Object {
	return obj.SendMsg("release")
}

// Send the "copy" message to the Object.
func (obj *Object) Copy() *Object {
	return obj.SendMsg("copy")
}

// Return representation of the Object suitable for printing.
// Under the hood, this method calls "description" on the Object.
func (obj *Object) String() string {
	pool := GetClass("NSAutoreleasePool").Alloc().Init()
	defer pool.Release()

	descString := obj.SendMsg("description")
	utf8Bytes := descString.SendMsg("UTF8String")
	if utf8Bytes != nil {
		return C.GoString((*C.char)(unsafe.Pointer(utf8Bytes.Pointer())))
	}

	return "(nil)"
}