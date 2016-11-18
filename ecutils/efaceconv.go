package ecutils

import "unsafe"

var strEface interface{}
var strKind uintptr
var sbEface interface{}
var sbKind uintptr

func init() {
	strEface = interface{}("")
	strKind = GetKind(strEface)

	sbEface = interface{}([]byte{})
	sbKind = GetKind(sbEface)
}

// GetKind returns arg's kind
// panics if arg is a pointer to value
func GetKind(arg interface{}) uintptr {
	return *(*uintptr)(unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&arg)))[0]))
}

// GetDataPtr returns pointer to arg's data
// panics if arg is a pointer to value
func GetDataPtr(arg interface{}) unsafe.Pointer {
	return unsafe.Pointer((*(*[2]uintptr)(unsafe.Pointer(&arg)))[1])
}

// Eface2String returns pointer to string and true if arg is a string
// or nil and false otherwise
func Eface2String(arg interface{}) (*string, bool) {
	if GetKind(arg) == strKind {
		return (*string)(GetDataPtr(arg)), true
	}
	return nil, false
}

// Eface2ByteSlice returns pointer to []byte and true if arg is a []byte
// or nil and false otherwise
func Eface2ByteSlice(arg interface{}) (*[]byte, bool) {
	if GetKind(arg) == sbKind {
		return (*[]byte)(GetDataPtr(arg)), true
	}
	return nil, false
}
