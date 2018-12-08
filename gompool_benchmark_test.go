package gompool

import (
	"bytes"
	"sync"
	"testing"
)

func BenchmarkGompool(b *testing.B) {
	pools := NewGompool(100, func() interface{} {
		return &bytes.Buffer{}
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool1 := pools.Get()
		_ = pool1.Value.(*bytes.Buffer)

		pool2 := pools.Get()
		_ = pool2.Value.(*bytes.Buffer)

		pools.Put(pool1)
		pools.Put(pool2)
	}
}

func BenchmarkDefaultPool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf1 := pool.Get().(*bytes.Buffer)

		buf2 := pool.Get().(*bytes.Buffer)

		pool.Put(buf1)
		pool.Put(buf2)
	}
}
