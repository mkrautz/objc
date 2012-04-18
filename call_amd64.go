package objc

/*
extern unsigned long GoObjc_CallTargetFrameSetup;
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type amd64frame struct {
	rdi  uintptr
	rsi  uintptr
	rdx  uintptr
	rcx  uintptr
	r8   uintptr
	r9   uintptr
	xmm0 float64
	xmm1 float64
	xmm2 float64
	xmm3 float64
	xmm4 float64
	xmm5 float64
	xmm6 float64
	xmm7 float64
}

func methodCallTarget() unsafe.Pointer {
	return unsafe.Pointer(&C.GoObjc_CallTargetFrameSetup)
}

//export goMethodCallEntryPoint
func goMethodCallEntryPoint(p uintptr) uintptr {
	frame := (*amd64frame)(unsafe.Pointer(p))

	obj := object{ptr: frame.rdi}
	sel := selectorToString(frame.rsi)
	clsName := object{ptr: obj.SendMsg("class").Pointer()}.className()

	clsInfo := classMap[clsName]
	method := clsInfo.MethodForSelector(sel)

	ptr := obj.internalPointer()
	selfVal := reflect.NewAt(clsInfo.typ, ptr)
	methodVal := reflect.ValueOf(method)

	args := []reflect.Value{selfVal, reflect.ValueOf(obj)}
	methodVal.Call(args)

	return 0
}
