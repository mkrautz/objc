// Package objc implements access to the Objective-C runtime from Go
package objc

/*
#cgo LDFLAGS: -lobjc -framework Foundation
#include <objc/runtime.h>
#include <objc/message.h>
#include <stdio.h>
#include <math.h>

void *GoObjc_GetClass(char *name) {
	return (void *) objc_getClass(name);
}

void *GoObjc_AllocateClassPair(void *superCls, char *name) {
	void *cls = objc_allocateClassPair(superCls, name, 0);
	if (class_addIvar(cls, "__go_internal", sizeof(void *), log2(sizeof(void *)), "^") == YES)
		return cls;
	return NULL;
}

void GoObjc_ClassAddMethod(void *subCls, void *sel, void *imp, char *typ) {
	class_addMethod(subCls, sel, imp, typ);
}

void GoObjc_SetInternal(void *obj, void *cls, void *ptr) {
	Ivar iv = class_getInstanceVariable(cls, "__go_internal");
	unsigned long *v = obj + ivar_getOffset(iv);
	*v = (unsigned long) ptr;
}

 void *GoObjc_GetInternal(void *obj, void *cls) {
	Ivar iv = class_getInstanceVariable(cls, "__go_internal");
	unsigned long *v = obj + ivar_getOffset(iv);
	return (void *) *v;
 }

void GoObjc_RegisterClass(void *cls) {
	objc_registerClassPair(cls);
}
*/
import "C"
import (
	"unsafe"
)

// An Object represents an Objective-C object, along with
// some convenience methods only found on NSObjects.
type Object interface {
	// Pointer returns the in-memory address of the object.
	Pointer() uintptr

	// SendMsg sends an arbitrary message to the method on the
	// object that is idenfieid by selectorName.
	SendMsg(selectorName string, args ...interface{}) Object

	// Alloc sends the  "alloc" message to the object.
	Alloc() Object

	// Init sends the "init" message to the object.
	Init() Object

	// Retain sends the "retain" message to the object.
	Retain() Object

	// Release sends the "release" message to the object.
	Release() Object

	// AutoRelease sends the "autorelease" message to the object.
	AutoRelease() Object

	// Copy sends the "copy" message to the object.
	Copy() Object

	// String returns a string-representation of the object.
	// This is equivalent to sending the "description"
	// message to the object, except that this method
	// returns a Go string.
	String() string
}

// Type object is the package's internal representation of an Object.
// Besides implementing the Objct interface, object also implements
// the Class interface.
type object uintptr

// Return the Object as a uintptr.
//
// Using package unsafe, this uintptr can further
// be converted to an unsafe.Pointer.
func (obj object) Pointer() uintptr {
	return uintptr(unsafe.Pointer(obj))
}

func (obj object) Alloc() Object {
	return obj.SendMsg("alloc")
}

func (obj object) Init() Object {
	return obj.SendMsg("init")
}

func (obj object) Retain() Object {
	return obj.SendMsg("retain")
}

func (obj object) Release() Object {
	return obj.SendMsg("release")
}

func (obj object) AutoRelease() Object {
	return obj.SendMsg("autorelease")
}

func (obj object) Copy() Object {
	return obj.SendMsg("copy")
}

func (obj object) String() string {
	pool := GetClass("NSAutoreleasePool").Alloc().Init()
	defer pool.Release()

	descString := obj.SendMsg("description")
	utf8Bytes := descString.SendMsg("UTF8String")
	if utf8Bytes != nil {
		return C.GoString((*C.char)(unsafe.Pointer(utf8Bytes.Pointer())))
	}

	return "(nil)"
}

// A Class represents a special Objective-C
// class Object.
type Class interface {
	Object
	AddMethod(selector string, typeInfo string)
}

// NewClass returns a new class that is a subclass of
// the specified superclass.
func NewClass(superClass Class, name string) Class {
	ptr := C.GoObjc_AllocateClassPair(unsafe.Pointer(superClass.Pointer()), C.CString(name))
	if ptr == nil {
		panic("unable to AllocateClassPair")
	}
	return object(uintptr(ptr))
}

// Lookup a Class by name
func GetClass(name string) Class {
	return object(C.GoObjc_GetClass(C.CString(name)))
}

// RegisterClass registers a Class with the Objective-C runtime.
func RegisterClass(class Class) {
	C.GoObjc_RegisterClass(unsafe.Pointer(class.Pointer()))
}

// AddMethod adds a new method to a Class.
func (cls object) AddMethod(selector string, typeInfo string) {
	sel := selectorWithName(selector)
	C.GoObjc_ClassAddMethod(unsafe.Pointer(cls.Pointer()), sel, methodCallTarget(), C.CString(typeInfo))
}

// setInternalPointer sets an internal pointer on the object.
// This is used to implement correct method dispatch for
// Objective-C classes created from within Go.
func (obj object) setInternalPointer(value unsafe.Pointer) {
	cls := obj.SendMsg("class")
	C.GoObjc_SetInternal(unsafe.Pointer(obj.Pointer()), unsafe.Pointer(cls.Pointer()), unsafe.Pointer(value))
}

// internalPointer returns the object's internal pointer.
// Must only be called on objects that are known to have
// an internal pointer set.
func (obj object) internalPointer() unsafe.Pointer {
	cls := obj.SendMsg("class")
	return C.GoObjc_GetInternal(unsafe.Pointer(obj.Pointer()), unsafe.Pointer(cls.Pointer()))
}
