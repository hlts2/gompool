package lockfree

import (
	"sync"
	"testing"
)

func BenchmarkThisLibrary(b *testing.B) {
	lf := New()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lf.Wait()
			lf.Signal()
		}
	})
}

func BenchmarkLockUnlock(b *testing.B) {
	m := new(sync.Mutex)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Lock()
			m.Unlock()
		}
	})
}
