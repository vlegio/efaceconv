# efaceconv
High performance conversion from interface{} to string without additional allocations
```
BenchmarkEface2String-4         100000000               10.2 ns/op             0 B/op          0 allocs/op
--- BENCH: BenchmarkEface2String-4
        efaceconv_test.go:29: string true
        efaceconv_test.go:29: string true
        efaceconv_test.go:29: string true
        efaceconv_test.go:29: string true
        efaceconv_test.go:29: string true
BenchmarkClassic-4              30000000                45.1 ns/op            16 B/op          1 allocs/op
--- BENCH: BenchmarkClassic-4
        efaceconv_test.go:43: string true
        efaceconv_test.go:43: string true
        efaceconv_test.go:43: string true
        efaceconv_test.go:43: string true
        efaceconv_test.go:43: string true
PASS
ok      github.com/t0pep0/efaceconv     2.442s
