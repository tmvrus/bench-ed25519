# bench-ed25519

go test -bench . -cpu 1,2,3,4  -benchtime=1s -benchmem

```
goos: darwin
goarch: amd64
pkg: github.com/tmvrus/bench-ed25519
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_Sign             36067             29604 ns/op              88 B/op          2 allocs/op
Benchmark_Sign-2           37072             29782 ns/op              88 B/op          2 allocs/op
Benchmark_Sign-3           38056             30789 ns/op              88 B/op          2 allocs/op
Benchmark_Sign-4           35690             29675 ns/op              88 B/op          2 allocs/op
Benchmark_Verify           19029             61950 ns/op               0 B/op          0 allocs/op
Benchmark_Verify-2         17564             62745 ns/op               0 B/op          0 allocs/op
Benchmark_Verify-3         18669             61759 ns/op               0 B/op          0 allocs/op
Benchmark_Verify-4         18308             62375 ns/op               0 B/op          0 allocs/op

```
