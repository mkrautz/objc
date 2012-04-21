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

var (
	objectInterfaceType   = reflect.TypeOf((*Object)(nil)).Elem()
	classInterfaceType    = reflect.TypeOf((*Class)(nil)).Elem()
	selectorInterfaceType = reflect.TypeOf((*Selector)(nil)).Elem()
)

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
