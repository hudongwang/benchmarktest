```shell
$ go test -bench='Add$' -cpu=2,4 -benchtime=5s -count=3 .
goos: darwin
goarch: arm64
pkg: github.com/hudongwang/benchmarktest/bmt001godemo
BenchmarkAdd-2          1000000000               0.3189 ns/op
BenchmarkAdd-2          1000000000               0.3156 ns/op
BenchmarkAdd-2          1000000000               0.3136 ns/op
BenchmarkAdd-4          1000000000               0.3132 ns/op
BenchmarkAdd-4          1000000000               0.3133 ns/op
BenchmarkAdd-4          1000000000               0.3133 ns/op
PASS
ok      github.com/hudongwang/benchmarktest/bmt001godemo        2.349s
```