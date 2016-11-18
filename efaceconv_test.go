package efaceconv

import (
	"testing"
)

var str = "string"

func TestEface2String(t *testing.T) {
	res, ok := Eface2String(str)
	if !ok {
		t.Error("Wrong type!")
	}
	if *res != str {
		t.Error("Not equal")
	}
	_, ok = Eface2String(ok)
	if ok {
		t.Error("Wrong type!")
	}
}

func BenchmarkEface2String(b *testing.B) {
	var v *string
	var ok bool
	for n := 0; n < b.N; n++ {
		v, ok = Eface2String(str)
	}
	b.Log(*v, ok)
}

func classic(arg interface{}) (v string, ok bool) {
	v, ok = arg.(string)
	return v, ok
}

func BenchmarkClassic(b *testing.B) {
	var v string
	var ok bool
	for n := 0; n < b.N; n++ {
		v, ok = classic(str)
	}
	b.Log(v, ok)
}
