package foundation

import (
	"log"
	"testing"
)

func TestString(t *testing.T) {
	str := NSStringFromString("hello, world!")
	log.Printf("%v", str)
}