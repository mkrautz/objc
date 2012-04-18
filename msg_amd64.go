package objc

import (
	"github.com/mkrautz/variadic"
	"math"
	"reflect"
	"unsafe"
)

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

func (obj object) SendMsg(selector string, args ...interface{}) Object {
	// Keep ObjC semantics: messages can be sent to nil objects,
	// but the response is nil.
	if obj.ptr == 0 {
		return nil
	}

	sel := selectorWithName(selector)
	if sel == nil {
		return nil
	}

	intArgs := []uintptr{}
	floatArgs := []uintptr{}
	memArgs := []uintptr{}

	var typeInfo string

	for i, arg := range args {
		switch t := arg.(type) {
		case Object:
			intArgs = append(intArgs, t.Pointer())
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
		// Float64 is a bit of a special case. Since SendMsg is a variadic
		// Go function, implicit floats will be of type float64, but we can't
		// be sure that the receiver expects that; they might expect a float32
		// instead.
		//
		// To remedy this, we query the selector's type encoding, and check
		// whether it expects a 32-bit or 64-bit float.	
		case float64:
			// Request typeInfo if we don't have it already.
			if typeInfo == "" {
				typeInfo = simpleTypeInfoForMethod(obj, selector)
			}
			typeEnc := string(typeInfo[i+3])
			switch typeEnc {
			case encFloat:
				floatArgs = append(floatArgs, uintptr(math.Float32bits(float32(t))))
			case encDouble:
				floatArgs = append(floatArgs, uintptr(math.Float64bits(t)))
			default:
				panic("objc: float argument mismatch")
			}
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

	fc.NumFloat = int64(len(floatArgs))
	for i, v := range floatArgs {
		fc.Words[6+i] = v
	}

	return object{ptr: fc.Call()}
}
