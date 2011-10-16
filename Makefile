include $(GOROOT)/src/Make.inc

TARG=github.com/mkrautz/go-objc

CGOFILES=\
	objc.go

include $(GOROOT)/src/Make.pkg
