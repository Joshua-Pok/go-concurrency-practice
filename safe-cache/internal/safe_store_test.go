package internal

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestEnvironment(t *testing.T) {
	t.Log("safety protocol intialized")
}

func TestConcurrentReadWrite(t *testing.T) {

	store := NewSafeStore[int]()
	store.Set("first", 1)

	var wg sync.WaitGroup

	//writers
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			store.Set("first", i)

		}()
	}

	//readers
	for i := 0; i < 500; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			store.Get("first")
		}()

	}

	wg.Wait() //wait until all done
}

func BenchmarkConcurrentCache(b *testing.B) {

	store := NewSafeStore[int]()
	var counter atomic.Int64
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := "first"

			if counter.Add(1)%10 == 0 {
				store.Set(key, 42)
			} else {
				store.Get(key)
			}
		}
	})
}
