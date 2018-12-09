# lock-free
lock-free is simple fastest lock-free library based on [cas](https://en.wikipedia.org/wiki/Compare-and-swap) written in golang.

## Requirement

Go (>=1.8)

## Installation

```shell
go get github.com/hlts2/lock-free
```

## Example

```go
wg := new(sync.WaitGroup)

lf := lockfree.New()

for i := 0; i < size; i++ {
    wg.Add(1)

    go func(i int) {
        defer wg.Done()

        // In the block between Wait and Signal, it becomes gruoute-safe
        lf.Wait()
        cnt++

        lf.Signal()
    }(i)
}

wg.Wait()

```

## Author
[hlts2](https://github.com/hlts2)

## LICENSE
lock-free released under MIT license, refer [LICENSE](https://github.com/hlts2/lock-free/blob/master/LICENSE) file.
