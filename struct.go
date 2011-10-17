package objc

// A type implementing the Struct interface can be
// passed directly as an argument to a SendMsg call
// to an Object.
type Struct interface {
	PackStruct64() []uint64
	PackStruct32() []uint32
}