package oid

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGoID(t *testing.T) {
	assert.Equal(t, ID(), ID())
}

func TestConcurrentGetID(t *testing.T) {
	var wg sync.WaitGroup
	concurrent := 1000
	wg.Add(concurrent)
	otherID := ID()
	for i := 0; i < concurrent; i++ {
		go func() {
			assert.Equal(t, ID(), ID())
			assert.NotEqual(t, otherID, ID())
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkGoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ID()
	}
}
