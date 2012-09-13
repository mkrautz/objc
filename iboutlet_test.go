// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"reflect"
	"testing"
	"unsafe"
)

type IBOutletTester struct {
	Object

	Myself Object `objc:"IBOutlet"`
}

const NSUTF8StringEncoding = 4

func NSStringFromString(str string) Object {
	hdrp := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return GetClass("NSString").SendMsg("alloc").SendMsg("initWithBytes:length:encoding:", hdrp.Data, hdrp.Len, NSUTF8StringEncoding)
}

func TestKeyValueCodingImpl(t *testing.T) {
	c := NewClass(GetClass("NSObject"), "IBOutletTester", IBOutletTester{})
	RegisterClass(c)

	ibo := new(IBOutletTester)
	NewGoInstance("IBOutletTester", ibo)

	ibo.SendMsg("setValue:forKey:", ibo, NSStringFromString("Myself").AutoRelease())

	if ibo.Myself == nil {
		t.Fatal("nil iboutlet, value not properly set for key")
	}

	if ibo.Myself.Pointer() != ibo.Pointer() {
		t.Error("value not set, or incorrectly set. myself=%p, struct=%p", ibo)
	}
}
