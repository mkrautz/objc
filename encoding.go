package objc

import (
	"reflect"
)

const (
	encId          = "@"
	encClass       = "#"
	encSelector    = ":"
	encChar        = "c"
	encUChar       = "C"
	encShort       = "s"
	encUShort      = "S"
	encInt         = "i"
	encUInt        = "I"
	encLong        = "l"
	encULong       = "L"
	encLongLong    = "q"
	encULongLong   = "Q"
	encFloat       = "f"
	encDouble      = "d"
	encDFLD        = "b"
	encBool        = "B"
	encVoid        = "v"
	encUndef       = "?"
	encPtr         = "^"
	encCharPtr     = "*"
	encAtom        = "%"
	encArrayBegin  = "["
	encArrayEnd    = "]"
	encUnionBegin  = "("
	encUnionEnd    = ")"
	encStructBegin = "{"
	encStructEnd   = "}"
	encVector      = "!"
	encConst       = "r"
)

func typeInfoForType(typ reflect.Type) string {
	kind := typ.Kind()
	switch kind {
	case reflect.Bool:
		return encBool
	case reflect.Int:
		return encInt
	case reflect.Int8:
		return encChar
	case reflect.Int16:
		return encShort
	case reflect.Int32:
		return encInt
	case reflect.Int64:
		return encULong
	case reflect.Uint:
		return encUInt
	case reflect.Uint8:
		return encUChar
	case reflect.Uint16:
		return encUShort
	case reflect.Uint32:
		return encUInt
	case reflect.Uint64:
		return encULong
	case reflect.Uintptr:
		return encPtr
	case reflect.Float32:
		return encFloat
	case reflect.Float64:
		return encDouble
	case reflect.Complex64, reflect.Complex128:
		// skip
	case reflect.Array:
		//skip
	case reflect.Chan:
		// skip
	case reflect.Func:
		// skip
	case reflect.Interface:
		return encId
	case reflect.Map:
		// skip
	case reflect.Ptr:
		// skip
	case reflect.Slice:
		// skip
	case reflect.String:
		// skip
	case reflect.Struct:
		// skip
	case reflect.UnsafePointer:
		// skip
	}
	return encPtr
}

// Returns the function's typeInfo
func funcTypeInfo(fn interface{}) string {
	typ := reflect.TypeOf(fn)
	kind := typ.Kind()
	if kind != reflect.Func {
		panic("not a func")
	}

	typeInfo := ""
	numOut := typ.NumOut()
	switch numOut {
	case 0:
		typeInfo += encVoid
	case 1:
		typeInfo += typeInfoForType(typ.Out(0))
	default:
		panic("too many output parameters")
	}

	for i := 0; i < typ.NumIn(); i++ {
		typeInfo += typeInfoForType(typ.In(i))
	}
	return typeInfo
}
