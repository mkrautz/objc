include $(GOROOT)/src/Make.inc

TARG=github.com/mkrautz/go-objc

CGOFILES=\
	objc.go\

GOFILES=\
	struct.go

include $(GOROOT)/src/Make.pkg
