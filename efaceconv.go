package efaceconv

import (
	"unsafe"
)

var strEface interface{}
var strKind uintptr
var sbEface interface{}
var sbKind uintptr

func init() {
	strEface = interface{}("")
	strKind = *(*uintptr)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&strEface)))[0]))
	sbEface = interface{}([]byte{})
	sbKind = *(*uintptr)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&sbEface)))[0]))
}

func Eface2String(arg interface{}) (*string, bool) {
	if *(*uintptr)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&arg)))[0])) == strKind {
		return (*string)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&arg)))[1])), true
	}
	return nil, false
}

func Eface2ByteSlice(arg interface{}) (*[]byte, bool) {
	if *(*uintptr)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&arg)))[0])) == sbKind {
		return (*[]byte)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&arg)))[1])), true
	}
	return nil, false
}
