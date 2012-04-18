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
