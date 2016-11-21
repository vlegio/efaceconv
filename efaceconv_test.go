package main

import (
	"testing"
)

const (
	ecImport         = "net/http"
	ecVar            = "_StringKind uintptr"
	ecInit           = "var sString string\n  _StringKind = ecutils.GetKind(sString)"
	ecMethodTemplate = `
// Eface2String returns pointer to string and true if arg is a string
// or nil and false otherwise
func Eface2String(arg interface{}) (*string, bool) {
	if ecutils.GetKind(arg) == _StringKind {
		return (*string)(ecutils.GetDataPtr(arg)), true
	}
	return nil, false
}

`
	ecTestTemplate = `
func TestEface2String(t *testing.T) {
  var String string
	res, ok := Eface2String(String)
	if !ok {
		t.Error("Wrong type!")
	}
	if !reflect.DeepEqual(*res, String) {
		t.Error("Not equal")
	}
	_, ok = Eface2String(ok)
	if ok {
		t.Error("Wrong type!")
	}
}

`

	ecBenchmark = `
func BenchmarkEface2String(b *testing.B) {
  var String string
	var v *string
	var ok bool
	for n := 0; n < b.N; n++ {
		v, ok = Eface2String(String)
	}
	b.Log(v, ok) //For don't use compiler optimization
}

func _StringClassic(arg interface{}) (v string, ok bool) {
	v, ok = arg.(string)
	return v, ok
}

func BenchmarkStringClassic(b *testing.B) {
  var String string
  var v string
	var ok bool
	for n := 0; n < b.N; n++ {
		v, ok = _StringClassic(String)
	}
	b.Log(v, ok) //For don't use compiler optimization
}


`
)

var unit = ecUnit{Import: "net/http", Type: "string", CustomName: "String"}

func TestEcUnitImport(t *testing.T) {
	if unit.generateImport() != ecImport {
		t.Fatal("Import not correct")
	}
}

func TestEcUnitInit(t *testing.T) {
	if unit.generateInit() != ecInit {
		t.Fatal("Init not correct")
	}
}

func TestEcUnitConvMethod(t *testing.T) {
	if unit.generateConvMethod() != ecMethodTemplate {
		t.Fatal("Method not correct")
	}
}

func TestEcUnitTestMethod(t *testing.T) {
	if unit.generateTestMethod() != ecTestTemplate {
		t.Fatal("Test method not correct")
	}
}

func TestEcUnitBenchmarkMethod(t *testing.T) {
	if unit.generatebenchmarkMethod() != ecBenchmark {
		t.Fatal("Benchmark not correct")
	}
}
