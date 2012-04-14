package appkit

import (
	"github.com/mkrautz/objc"
)

type NSImageView struct {
	*objc.Object
}

func NewNSImageView() NSImageView {
	return NSImageView{objc.GetClass("NSImageView").Alloc().Init()}
}

func (imgView NSImageView) SetImage(img NSImage) {
	imgView.SendMsg("setImage:", img.Object)
}

func (imgView NSImageView) Image() NSImage {
	return NSImage{imgView.SendMsg("image")}
}