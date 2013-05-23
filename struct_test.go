// Copyright (c) 2013 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"testing"
)

type NSPoint struct {
	X float32
	Y float32
}

type NSSize struct {
	Width  float32
	Height float32
}

type NSRect struct {
	Origin NSPoint
	Size   NSSize
}

func TestStructPassing(t *testing.T) {
	rect := NSRect{
		NSPoint{
			X: 0,
			Y: 0,
		},
		NSSize{
			Width: 100,
			Height: 100,
		},
	}

	obj := GetClass("NSWindow").Alloc().
			SendMsg("initWithContentRect:styleMask:backing:defer:",
				rect,  // rect
				0,     // windowStyle = default
				2,     // bufferingType = NSBackingStorebuffered
				false) // deferCreation
	if obj.Pointer() == 0 {
		t.Fatalf("unable to create NSWindow, got nil ptr")
	}
}