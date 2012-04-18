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

func TestDoubleReturnValue(t *testing.T) {
	in := float64(54.0)
	out := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithDouble:", in).SendMsg("doubleValue")
	if out.Float() != in {
		t.Errorf("expected %v, got %v", in, out.Float())
	}
}

func TestFloatReturnValue(t *testing.T) {
	in := float64(54.0)
	out := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithDouble:", in).SendMsg("floatValue")
	if out.Float() != in {
		t.Errorf("expected %v, got %v", in, out.Float())
	}
}
