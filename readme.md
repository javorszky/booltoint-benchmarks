# bool 2 int

This is a small repository listing a bunch of methods of converting a bool (`true`, `false`) into the relevant 
integers: (`0`, `1`).

It also adds benchmarks.

The source is this article: https://dev.to/chigbeef_77/bool-int-but-stupid-in-go-3jb3

## The fastest usable conversion

```go
func bool2intFastest(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue 6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}
```

Fun little article about compiler optimization and why this is the best: https://0x0f.me/blog/golang-compiler-optimization/.

The other one is using the unsafe pointer, much like the inverse square root hack, but I'm not about to introduce 
unsafe pointers into anything without a huge performance gain. The difference between the above and the unsafe 
pointer version is negligible.

## Benchmarks

```
$ make bench
go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/javorszky/bool2init
cpu: Apple M1 Pro
Benchmark_bool2intSimplest-10           362035316                3.304 ns/op           0 B/op          0 allocs/op
Benchmark_bool2intFastest-10            525728955                2.397 ns/op           0 B/op          0 allocs/op
Benchmark_bool2intSwitch-10             367150812                3.251 ns/op           0 B/op          0 allocs/op
Benchmark_bool2intInternalMap-10        17308472                72.98 ns/op            0 B/op          0 allocs/op
Benchmark_bool2intExternalMap-10        72333274                16.59 ns/op            0 B/op          0 allocs/op
Benchmark_bool2intSlice-10              307067509                3.900 ns/op           0 B/op          0 allocs/op
Benchmark_bool2intUnsafePointer-10      527730302                2.275 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/javorszky/bool2init  8.768s
```