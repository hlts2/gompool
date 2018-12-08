package gompool

import (
	"testing"

	"github.com/hlts2/gompool/treiber"
)

func TestNewGompool(t *testing.T) {
	pool := NewGompool(12, func() interface{} {
		return new(string)
	})

	if pool == nil {
		t.Errorf("NewGompool is nil")
	}
}

func TestGetAndPut(t *testing.T) {
	var poolSize uint = 100
	pool := NewGompool(poolSize, func() interface{} {
		return make([]int, 10)
	})

	poolNodes := make([]*treiber.Node, 0, int(poolSize))

	for i := 0; i < int(poolSize); i++ {
		n, err := pool.Get()
		if err != nil {
			t.Errorf("Get is error: %v", err)
		}

		poolNodes = append(poolNodes, n)
	}

	expected := 0
	got := pool.Cap()

	if expected != got {
		t.Errorf("Cap expected: %v, got: %v", expected, got)
	}

	for _, n := range poolNodes {
		pool.Put(n)
	}

	expected = int(poolSize)
	got = pool.Cap()

	if expected != got {
		t.Errorf("Cap expected: %v, got: %v", expected, got)
	}
}

func TestCap(t *testing.T) {
	var poolSize uint = 10

	pool := NewGompool(poolSize, func() interface{} {
		return new(int)
	})

	got := pool.Cap()

	if got != int(poolSize) {
		t.Errorf("Cap expected: %v, got: %v", poolSize, got)
	}

	_, _ = pool.Get()

	got = pool.Cap()

	if got != int(poolSize)-1 {
		t.Errorf("Cap expected: %v, got: %v", int(poolSize)-1, got)
	}
}

func TestIsEmpty(t *testing.T) {
	var poolSize uint = 4

	pool := NewGompool(poolSize, func() interface{} {
		return new(int)
	})

	expected := false
	got := pool.IsEmpty()

	if expected != got {
		t.Errorf("IsEmpty expected: %v, got: %v", expected, got)
	}

	_, _ = pool.Get()
	_, _ = pool.Get()
	_, _ = pool.Get()
	_, _ = pool.Get()

	expected = true
	got = pool.IsEmpty()

	if expected != got {
		t.Errorf("IsEmpty expected: %v, got: %v", expected, got)
	}
}

func TestDestPool(t *testing.T) {
	var poolSize uint = 20

	pool := NewGompool(poolSize, func() interface{} {
		return make([]byte, 100)
	})

	expected := int(poolSize)
	got := pool.Cap()

	if expected != got {
		t.Errorf("Cap expected: %v, got: %v", expected, got)
	}

	pool.DestPool()

	expected = 0
	got = pool.Cap()

	if expected != got {
		t.Errorf("Cap expected: %v, got: %v", expected, got)
	}
}
