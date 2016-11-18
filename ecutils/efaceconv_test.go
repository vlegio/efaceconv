package ecutils

import (
	"testing"
)

var (
	str = "string"
	sb  = []byte("slice of byte")
)

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

func TestEface2ByteSlice(t *testing.T) {
	res, ok := Eface2ByteSlice(sb)
	if !ok {
		t.Error("Wrong type!")
	}
	if len(*res) != len(sb) {
		t.Error("Not equal")
	}
	for i := range *res {
		if (*res)[i] != sb[i] {
			t.Error("Not equal")
		}
	}
	_, ok = Eface2ByteSlice(ok)
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

func BenchmarkEface2ByteSlice(b *testing.B) {
	var v *[]byte
	var ok bool
	for n := 0; n < b.N; n++ {
		v, ok = Eface2ByteSlice(sb)
	}
	b.Log(*v, ok)
}

func sbClassic(arg interface{}) (v []byte, ok bool) {
	v, ok = arg.([]byte)
	return v, ok
}

func BenchmarkSBClassic(b *testing.B) {
	var v []byte
	var ok bool
	for n := 0; n < b.N; n++ {
		v, ok = sbClassic(sb)
	}
	b.Log(v, ok)
}
