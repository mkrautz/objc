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
	Object `objc:"IBOutletTester : NSObject"`
	Myself Object `objc:"IBOutlet"`
}

func (ibo *IBOutletTester) MyselfIsNil() bool {
	return ibo.Myself == nil
}

func (ibo *IBOutletTester) MyselfIsMyself() bool {
	return ibo.Myself.Pointer() == ibo.Object.Pointer()
}

const NSUTF8StringEncoding = 4

func NSStringFromString(str string) Object {
	hdrp := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return GetClass("NSString").SendMsg("alloc").SendMsg("initWithBytes:length:encoding:", hdrp.Data, hdrp.Len, NSUTF8StringEncoding)
}

func TestKeyValueCodingImpl(t *testing.T) {
	c := NewClass(IBOutletTester{})
	c.AddMethod("myselfIsNil", (*IBOutletTester).MyselfIsNil)
	c.AddMethod("myselfIsMyself", (*IBOutletTester).MyselfIsMyself)
	RegisterClass(c)

	pool := GetClass("NSAutoreleasePool").SendMsg("alloc").SendMsg("init")
	defer pool.SendMsg("release")

	ibo := GetClass("IBOutletTester").SendMsg("alloc").SendMsg("init")
	ibo.SendMsg("setValue:forKey:", ibo, NSStringFromString("Myself").AutoRelease())

	if ibo.SendMsg("myselfIsNil").Bool() {
		t.Fatal("nil iboutlet, value not properly set for key")
	}

	if !ibo.SendMsg("myselfIsMyself").Bool() {
		t.Error("value not set, or incorrectly set.")
	}
}
