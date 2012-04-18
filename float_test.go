package objc

import "testing"

func TestFloatArgsImplicit(t *testing.T) {
	expected := 54.0
	number := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithFloat:", expected)
	str := number.String()
	if str != "54" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestDoubleArgsImplicit(t *testing.T) {
	expected := 54.0
	number := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithDouble:", expected)
	str := number.String()
	if str != "54" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestFloatArgsExplicit(t *testing.T) {
	expected := float32(54.0)
	number := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithFloat:", expected)
	str := number.String()
	if str != "54" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestDoubleArgsExplicit(t *testing.T) {
	expected := float64(54.0)
	number := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithDouble:", expected)
	str := number.String()
	if str != "54" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}
