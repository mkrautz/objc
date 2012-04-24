package foundation

import (
	"github.com/mkrautz/objc"
)

type NSDictionary struct {
	objc.Object
}

func (dict NSDictionary) ObjectForKey(key objc.Object) objc.Object {
	return dict.SendMsg("objectForKey:", key)
}