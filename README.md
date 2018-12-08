# gompool

gompool is a simple lock-free memory pool library written in golang using [`treiber stack`][treiber stack]

[treiber stack]: https://en.wikipedia.org/wiki/Treiber_Stack

## Requirement

- go (>= 1.8)

## Installation

```shell
go get github.com/hlts2/gompool
```

## Example

```go

// Create 10 pools of `*bytes.Buffer`
pools := gompool.NewGompool(10, func() interface{} {
    return &bytes.Buffer{}
})

pool1 := pools.Get()
pool2 := pools.Get()

buf1 := pool1.Value.(*bytes.Buffer)
buf1.WriteString("Hello world1")

buf2 := pool2.Value.(*bytes.Buffer)
buf2.WriteString("Hello world2")

pools.Put(pool1)    // Return pool1 to pools
pools.Put(pool2)    // Return pool2 to pools

```

## Benchmark
[gompool](https://github.com/hlts2/gompool) vs [sync.Pool](https://github.com/golang/go/tree/master/src/sync)

```
goos: darwin
goarch: amd64
pkg: github.com/hlts2/gompool
BenchmarkGompool-4       	30000000	        54.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGompool-4       	30000000	        54.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGompool-4       	30000000	        55.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGompool-4       	30000000	        54.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGompool-4       	30000000	        53.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkDefaultPool-4   	20000000	        66.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkDefaultPool-4   	20000000	        67.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkDefaultPool-4   	20000000	        67.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkDefaultPool-4   	20000000	        68.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkDefaultPool-4   	20000000	        68.1 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/hlts2/gompool	15.530s
```

## Author
[hlts2](https://github.com/hlts2)

## LICENSE
lock-free released under MIT license, refer [LICENSE](https://github.com/hlts2/gompool/blob/master/LICENSE) file.
