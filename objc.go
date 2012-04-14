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
*/
import "C"
import (
	"github.com/mkrautz/variadic"
	"math"
	"reflect"
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

func unpackStruct(val reflect.Value) []uintptr {
	memArgs := []uintptr{}
	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i)
		kind := v.Kind()
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			memArgs = append(memArgs, uintptr(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			memArgs = append(memArgs, uintptr(v.Uint()))
		case reflect.Float32, reflect.Float64:
			memArgs = append(memArgs, uintptr(math.Float64bits(v.Float())))
		case reflect.Ptr:
			memArgs = append(memArgs, val.Pointer())
		case reflect.Struct:
			args := unpackStruct(v)
			memArgs = append(memArgs, args...)
		}
	}
	return memArgs
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

	intArgs := []uintptr{}
	floatArgs := []uintptr{}
	memArgs := []uintptr{}

	for i, arg := range args {
		switch t := arg.(type) {
		case uintptr:
			intArgs = append(intArgs, t)
		case int:
			intArgs = append(intArgs, uintptr(t))
		case uint:
			intArgs = append(intArgs, uintptr(t))
		case int8:
			intArgs = append(intArgs, uintptr(t))
		case uint8:
			intArgs = append(intArgs, uintptr(t))
		case int16:
			intArgs = append(intArgs, uintptr(t))
		case uint16:
			intArgs = append(intArgs, uintptr(t))
		case int32:
			intArgs = append(intArgs, uintptr(t))
		case uint32:
			intArgs = append(intArgs, uintptr(t))
		case int64:
			intArgs = append(intArgs, uintptr(t))
		case uint64:
			intArgs = append(intArgs, uintptr(t))
		case bool:
			if t {
				intArgs = append(intArgs, uintptr(1))
			} else {
				intArgs = append(intArgs, uintptr(0))
			}
		case float32:
			floatArgs = append(floatArgs, uintptr(math.Float32bits(t)))
		case float64:
			floatArgs = append(floatArgs, uintptr(math.Float64bits(t)))
		default:
			val := reflect.ValueOf(args[i])
			switch val.Kind() {
			case reflect.Ptr:
				intArgs = append(intArgs, val.Pointer())
			case reflect.Uintptr:
				intArgs = append(intArgs, uintptr(val.Uint()))
			case reflect.Struct:
				args := unpackStruct(val)
				memArgs = append(memArgs, args...)
			default:
				panic("unhandled kind")
			}
		}
	}

	fc := variadic.NewFunctionCall("objc_msgSend")
	fc.Words[0] = obj.Pointer()
	fc.Words[1] = uintptr(sel)

	if len(memArgs) > 0 {
		fc.Memory = unsafe.Pointer(&memArgs[0])
		fc.NumMemory = int64(len(memArgs))
	}

	if len(intArgs) > 4 {
		panic("too many int args")
	}
	if len(floatArgs) > 8 {
		panic("too many float args")
	}

	for i, v := range intArgs {
		fc.Words[i+2] = v
	}

	for i, v := range floatArgs {
		fc.Words[6+i] = v
	}

	return (*Object)(unsafe.Pointer(fc.Call()))
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
