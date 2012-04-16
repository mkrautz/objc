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