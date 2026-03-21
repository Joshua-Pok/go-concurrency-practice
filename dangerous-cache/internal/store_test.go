package main

import (
	"sync"
	"testing"
)

func TestSmoke(t *testing.T) {
	if false {
		t.Error("The Laws of physics has failed")
	}

}

func TestBasicOperations(t *testing.T) {
	store := NewStore[int]()

	store.Set("first", 1)

	res, ok := store.Get("first")
	if !ok {
		t.Errorf("couldnt get key")
	}

	if res != 1 {
		t.Errorf("wrong value associated with key")
	}
}

func TestConcurrentReadWrite(t *testing.T) {

	store := NewStore[int]()
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
