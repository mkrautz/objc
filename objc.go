// Package objc implements access to the Objective-C runtime from Go
package objc

/*
#cgo LDFLAGS: -lobjc -framework Foundation
#include <objc/runtime.h>
#include <objc/message.h>

void *GoObjc_GetClass(char *name) {
	return (void *) objc_getClass(name);
}

void *GoObjc_RegisterSelector(char *name) {
	return (void *) sel_registerName(name);
}

void *GoObjc_AllocateClassPair(void *superCls, char *name) {
	return (void *) objc_allocateClassPair(superCls, name, 0);
}

void GoObjc_ClassAddMethod(void *subCls, void *sel, void *imp, char *typ) {
	class_addMethod(subCls, sel, imp, typ);
}

void GoObjc_RegisterClass(void *cls) {
	objc_registerClassPair(cls);
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

// NewClass returns a new class that is a sublcass of
// the specified super class.
func NewClass(superClass *Object, name string) *Object {
	return (*Object)(C.GoObjc_AllocateClassPair(unsafe.Pointer(superClass), C.CString(name)))
}

// RegisterClass registers an object representing a class
// with the Objective-C runtine.
func RegisterClass(class *Object) {
	C.GoObjc_RegisterClass(unsafe.Pointer(class))	
}

// AddMethod adds a new method to a class.
func (cls *Object) AddMethod(s Selector, typeInfo string) {
	C.GoObjc_ClassAddMethod(unsafe.Pointer(cls), unsafe.Pointer(uintptr(s)), methodCallTarget(), C.CString(typeInfo))
}

// An Object represents an Objective-C object, but it also implements convenience
// methods represent methods usually found on Foundation's NSObject class.
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
	return obj.SendMsg("autorelease")
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
