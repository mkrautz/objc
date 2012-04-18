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
	retVals := methodVal.Call(args)

	if len(retVals) > 0 {
		val := retVals[0]
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return uintptr(val.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return uintptr(val.Uint())
		case reflect.Bool:
			if val.Bool() {
				return 1
			} else {
				return 0
			}
		case reflect.Float32, reflect.Float64:
			panic("objc: float return values not yet supported")
		case reflect.Interface:
			if obj, ok := val.Interface().(Object); ok {
				return obj.Pointer()
			}
			panic("objc: bad interface return value")
		default:
			panic("objc: unknown return value")
		}
	}

	return 0
}
