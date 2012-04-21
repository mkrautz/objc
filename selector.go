package objc

/*
#cgo LDFLAGS: -lobjc
#include <objc/runtime.h>

void *GoObjc_RegisterSelector(char *name) {
	return (void *) sel_registerName(name);
}

char *GoObjc_TypeInfoForMethod(void *cls, void *sel) {
	Method m = class_getInstanceMethod(cls, sel);
	return (char *) method_getTypeEncoding(m);
}
*/
import "C"
import "unsafe"

// A Selector represents an Objective-C method selector.
type Selector interface {
	// Selector returns a string representation of
	// a selector.
	Selector() string

	// String returns the same string as Selector does.
	// It is only implemented to implement the Stringer
	// interface.
	String() string
}

// Type selector is the underlying implementation
// of the Selector interface. It is represented as
// a Go string.
type selector string

// Selector implements the Selector method of the
// Selector interface.
func (sel selector) Selector() string {
	return string(sel)
}

// String implements the String method of the
// Selector interface.
func (sel selector) String() string {
	return sel.Selector()
}

// GetSelector looks up a Selector by name.
func GetSelector(name string) Selector {
	return selector(name)
}

// selectorWithName looks up a selector by name.
func selectorWithName(name string) unsafe.Pointer {
	return C.GoObjc_RegisterSelector(C.CString(name))
}

// typeInfoForMethod returns the type encoding string for
// selector on obj's Class.
func typeInfoForMethod(obj Object, selector string) string {
	sel := selectorWithName(selector)
	cls := getObjectClass(obj)
	return C.GoString(C.GoObjc_TypeInfoForMethod(unsafe.Pointer(cls.Pointer()), sel))
}

// simpleTypeInfoForMethod returns a simplified typeInfo representation
// with C specifiers and stack information stripped out.
func simpleTypeInfoForMethod(obj Object, selector string) string {
	ti := typeInfoForMethod(obj, selector)
	sti := []byte{}
	for i := 0; i < len(ti); i++ {
		if ti[i] >= '0' && ti[i] <= '9' {
			continue
		}
		if string(ti[i]) == encConst {
			continue
		}
		// fixme(mkrautz): What is V? The NSObject release method uses V.
		if ti[i] == 'V' {
			continue
		}
		sti = append(sti, ti[i])
	}
	return string(sti)
}
