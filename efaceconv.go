package efaceconv

import (
	"unsafe"
)

var eface interface{}
var strKind uintptr

func init() {
	eface = interface{}("")
	strKind = *(*uintptr)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&eface)))[0]))
}

func Eface2String(arg interface{}) (*string, bool) {
	if *(*uintptr)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&arg)))[0])) == strKind {
		return (*string)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&arg)))[1])), true
	}
	return nil, false
}
