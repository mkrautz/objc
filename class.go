package objc

/*
#cgo LDFLAGS: -lobjc -framework Foundation
#include <objc/runtime.h>
#include <objc/message.h>
#include <stdio.h>
#include <math.h>

static unsigned long key = 0xbadc0c0a;

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

void GoObjc_SetInternal(void *obj, void *ptr) {
	objc_setAssociatedObject(obj, (void *)&key, ptr, OBJC_ASSOCIATION_ASSIGN);
}

void *GoObjc_GetInternal(void *obj) {
	return (void *) objc_getAssociatedObject(obj, (void *)&key);
}

void GoObjc_RegisterClass(void *cls) {
	objc_registerClassPair(cls);
}

char *GoObjc_GetClassName(void *cls) {
	return (char *) class_getName(cls);
}

char *GoObjc_SelectorToString(void *sel) {
	return (char *) sel;
}

*/
import "C"
import (
	"reflect"
	"unsafe"
)

type classInfo struct {
	typ       reflect.Type
	methodMap map[string]interface{}
}

func (ci classInfo) MethodForSelector(sel string) interface{} {
	return ci.methodMap[sel]
}

var (
	classMap map[string]classInfo
)

func init() {
	classMap = make(map[string]classInfo)
}

// A Class represents a special Objective-C
// class Object.
type Class interface {
	Object
	AddMethod(selector string, fn interface{})
}

// NewClass returns a new class, named by the name parameter,
// that is a subclass of the specified superclass.
//
// The value parameter must point to a value of the struct that
// is used to represent instances of the class in Go.
func NewClass(superClass Class, name string, value interface{}) Class {
	ptr := C.GoObjc_AllocateClassPair(unsafe.Pointer(superClass.Pointer()), C.CString(name))
	if ptr == nil {
		panic("unable to AllocateClassPair")
	}

	classMap[name] = classInfo{
		typ:       reflect.TypeOf(value),
		methodMap: make(map[string]interface{}),
	}

	return object{ptr: uintptr(ptr)}
}

// NewGoInstance registers a Go Obejctive-C class instance
// as an instance with the objc package.
//
// The className parameter must be a valid Objective-C class,
// typically one registered with the objc package by calling
// objc.NewClass.
//
// The value parameter must be a pointer to a struct that embeds
// the objc.Object interface.
func NewGoInstance(className string, value interface{}) {
	val := reflect.ValueOf(value)
	if val.Kind() != reflect.Ptr {
		panic("NewGoInstance: value must be a pointer")
	}

	// Extract the Object field of the embedded objc.Object,
	// if there is one.
	ptrval := reflect.ValueOf(value).Elem()
	var objval reflect.Value
	if ptrval.Kind() == reflect.Struct {
		objval = ptrval.FieldByName("Object")
		if !val.IsValid() {
			return
		}
	}

	// Set the Object field to a new instance of the class.
	o := GetClass(className).SendMsg("alloc").SendMsg("init")
	obj, ok := o.(Object)
	if !ok {
		panic("NewGoInstance: value does not implement Object")
	}
	objval.Set(reflect.ValueOf(obj))

	// Point the instance's internal pointer to the struct
	// that we're proxying our Objective-C object to.
	object{ptr: obj.Pointer()}.setInternalPointer(unsafe.Pointer(val.Pointer()))
}

// Lookup a Class by name
func GetClass(name string) Class {
	return object{ptr: uintptr(C.GoObjc_GetClass(C.CString(name)))}
}

// RegisterClass registers a Class with the Objective-C runtime.
func RegisterClass(class Class) {
	C.GoObjc_RegisterClass(unsafe.Pointer(class.Pointer()))
}

// className returns the name of the Class represented by object.
func (cls object) className() string {
	return C.GoString(C.GoObjc_GetClassName(unsafe.Pointer(cls.Pointer())))
}

// AddMethod adds a new method to a Class.
func (cls object) AddMethod(selector string, fn interface{}) {
	sel := selectorWithName(selector)
	typeInfo := funcTypeInfo(fn)
	C.GoObjc_ClassAddMethod(unsafe.Pointer(cls.Pointer()), sel, methodCallTarget(), C.CString(typeInfo))

	// Add the method to the class's method map
	clsName := cls.className()
	clsInfo := classMap[clsName]
	clsInfo.methodMap[selector] = fn
}

// setInternalPointer sets an internal pointer on the object.
// This is used to implement correct method dispatch for
// Objective-C classes created from within Go.
func (obj object) setInternalPointer(value unsafe.Pointer) {
	C.GoObjc_SetInternal(unsafe.Pointer(obj.Pointer()), unsafe.Pointer(value))
}

// internalPointer returns the object's internal pointer.
// Must only be called on objects that are known to have
// an internal pointer set.
func (obj object) internalPointer() unsafe.Pointer {
	return C.GoObjc_GetInternal(unsafe.Pointer(obj.Pointer()))
}

// selectorToString returns the string representation
// of the selector sel.
func selectorToString(sel uintptr) string {
	return C.GoString(C.GoObjc_SelectorToString(unsafe.Pointer(sel)))
}
