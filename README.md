[![TravisCI](https://travis-ci.org/t0pep0/efaceconv.svg?branch=master)](https://travis-ci.org/github.com/t0pep0/efaceconv)
[![GoReportCard](https://goreportcard.com/badge/github.com/t0pep0/efaceconv)](https://goreportcard.com/github.com/t0pep0/efaceconv)
[![GoDoc](https://godoc.org/github.com/t0pep0/efaceconv/ecutils?status.svg)](https://godoc.org/github.com/t0pep0/efaceconv)
[![Coverage Status](https://coveralls.io/repos/github/t0pep0/efaceconv/badge.svg?branch=master)](https://coveralls.io/github/t0pep0/efaceconv?branch=master)
# efaceconv
High performance conversion from interface{} to string without additional allocations

This is tool for go generate and common lib (ecutils)

Usage:

```go
\\go:generate efaceconv
\\ec::string:String
\\ec:net\http:http.Writer:HttpWritter
```

\\\\ec:\<import package name\>:\<type\>:\<custom name\>

generated function:
```go
func Eface2CustomName(arg interface{}) (*type, bool)
```

View  ./example dir

```
BenchmarkEface2String-4         100000000               10.8 ns/op             0 B/op          0 allocs/op
--- BENCH: BenchmarkEface2String-4
        efaceconv_test.go:51: string true
        efaceconv_test.go:51: string true
        efaceconv_test.go:51: string true
        efaceconv_test.go:51: string true
        efaceconv_test.go:51: string true
BenchmarkClassic-4              30000000                46.4 ns/op            16 B/op          1 allocs/op
--- BENCH: BenchmarkClassic-4
        efaceconv_test.go:66: string true
        efaceconv_test.go:66: string true
        efaceconv_test.go:66: string true
        efaceconv_test.go:66: string true
        efaceconv_test.go:66: string true
BenchmarkEface2ByteSlice-4      100000000               11.2 ns/op             0 B/op          0 allocs/op
--- BENCH: BenchmarkEface2ByteSlice-4
        efaceconv_test.go:75: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
        efaceconv_test.go:75: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
        efaceconv_test.go:75: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
        efaceconv_test.go:75: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
        efaceconv_test.go:75: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
BenchmarkSBClassic-4            30000000                56.9 ns/op            32 B/op          1 allocs/op
--- BENCH: BenchmarkSBClassic-4
        efaceconv_test.go:89: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
        efaceconv_test.go:89: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
        efaceconv_test.go:89: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
        efaceconv_test.go:89: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
        efaceconv_test.go:89: [115 108 105 99 101 32 111 102 32 98 121 116 101] true
PASS
ok      github.com/t0pep0/efaceconv     5.437s
```
