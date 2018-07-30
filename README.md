# gompool

gompool is a simple lock-free memory pool library in golang using [`treiber stack`][treiber stack]

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

pool1, _ := pools.Get()
pool2, _ := pools.Get()

buf1 := pool1.Value.(*bytes.Buffer)
buf1.WriteString("Hello world1")

buf2 := pool2.Value.(*bytes.Buffer)
buf2.WriteString("Hello world2")

pools.Put(pool1)    // Return pool1 to pools
pools.Put(pool2)    // Return pool2 to pools

```
