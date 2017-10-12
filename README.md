[![TravisCI](https://travis-ci.org/t0pep0/efaceconv.svg?branch=master)](https://travis-ci.org/t0pep0/efaceconv)
[![GoReportCard](https://goreportcard.com/badge/github.com/t0pep0/efaceconv)](https://goreportcard.com/report/github.com/t0pep0/efaceconv)
[![GoDoc](https://godoc.org/github.com/t0pep0/efaceconv/ecutils?status.svg)](https://godoc.org/github.com/t0pep0/efaceconv/ecutils)
[![Coverage Status](https://coveralls.io/repos/github/t0pep0/efaceconv/badge.svg?branch=master)](https://coveralls.io/github/t0pep0/efaceconv?branch=master)
# efaceconv
High performance conversion from interface{} to immutable types without additional allocations

This is tool for go generate and common lib (ecutils)

Usage:

```go
//go:generate efaceconv
//ec::string:String
//ec:net/http:http.ResponseWriter:ResponseWritter
```

//ec:\<import package name\>:\<type\>:\<custom name\>

generated function:
```go
func Eface2CustomName(arg interface{}) (*type, bool)
```

View  ./example dir

```
=== RUN   TestEface2SByte
--- PASS: TestEface2SByte (0.00s)
=== RUN   TestEface2String
--- PASS: TestEface2String (0.00s)
=== RUN   TestEface2SInt
--- PASS: TestEface2SInt (0.00s)
BenchmarkEface2SByte-4          100000000               11.8 ns/op             0 B/op          0 allocs/op
--- BENCH: BenchmarkEface2SByte-4
        efaceconv_generated_test.go:33: &[] true
        efaceconv_generated_test.go:33: &[] true
        efaceconv_generated_test.go:33: &[] true
        efaceconv_generated_test.go:33: &[] true
        efaceconv_generated_test.go:33: &[] true
BenchmarkSByteClassic-4         30000000                50.4 ns/op            32 B/op          1 allocs/op
--- BENCH: BenchmarkSByteClassic-4
        efaceconv_generated_test.go:48: [] true
        efaceconv_generated_test.go:48: [] true
        efaceconv_generated_test.go:48: [] true
        efaceconv_generated_test.go:48: [] true
        efaceconv_generated_test.go:48: [] true
BenchmarkEface2String-4         100000000               11.1 ns/op             0 B/op          0 allocs/op
--- BENCH: BenchmarkEface2String-4
        efaceconv_generated_test.go:76: 0xc42003fee8 true
        efaceconv_generated_test.go:76: 0xc420043ea8 true
        efaceconv_generated_test.go:76: 0xc420043ea8 true
        efaceconv_generated_test.go:76: 0xc420043ea8 true
        efaceconv_generated_test.go:76: 0xc420043ea8 true
BenchmarkStringClassic-4        30000000                45.3 ns/op            16 B/op          1 allocs/op
--- BENCH: BenchmarkStringClassic-4
        efaceconv_generated_test.go:91:  true
        efaceconv_generated_test.go:91:  true
        efaceconv_generated_test.go:91:  true
        efaceconv_generated_test.go:91:  true
        efaceconv_generated_test.go:91:  true
BenchmarkEface2SInt-4           100000000               11.6 ns/op             0 B/op          0 allocs/op
--- BENCH: BenchmarkEface2SInt-4
        efaceconv_generated_test.go:119: &[] true
        efaceconv_generated_test.go:119: &[] true
        efaceconv_generated_test.go:119: &[] true
        efaceconv_generated_test.go:119: &[] true
        efaceconv_generated_test.go:119: &[] true
BenchmarkSIntClassic-4          30000000                50.5 ns/op            32 B/op          1 allocs/op
--- BENCH: BenchmarkSIntClassic-4
        efaceconv_generated_test.go:134: [] true
        efaceconv_generated_test.go:134: [] true
        efaceconv_generated_test.go:134: [] true
        efaceconv_generated_test.go:134: [] true
        efaceconv_generated_test.go:134: [] true
PASS
ok      github.com/t0pep0/efaceconv/example     8.050s
```
