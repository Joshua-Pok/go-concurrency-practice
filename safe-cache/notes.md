<!--markdownlint-disable-->



# Step 1: Mutex Protected Store

Mutex acts as a single occupancy mechanism, only one goroutine can hold the key, everyone else has to wait



## Step Goal : Implement SafeStore[T any] struct that satisfies Cache Interface. It must contain sync.Mutex and protect every single access to internal map


Sketch


type SafeStore[T any] struct{
mu sync.RwMutex
data map[string]T
}


func (s *SafeStore) Set(key string, value T) error{
s.mu.Lock()
defer s.mu.Unlock() //make sure i dont forget to lock it


s.data[key] = value
return nil
}


func NewSafeStore[T any]() Cache[T]{
return &SafeStore[T]{
mu: sync.RWMutex
data: make(map[string]T)
}
}


func (s *SafeStore) Get(key string) (T, bool){
s.mu.RLock()

defer s.mu.RUnlock


//standard get code
}


func (s *SafeStore) Delete(key string) error{
s.mu.Lock()
defer s.mu.Unlock()


//standard delete code
}


# Step 2: RWMutex Refactor


RW mutex is a more advanced lock


**RLock() and RUnlock()** for readers. Multiple Go routines can hold a lock simultaenously. Lock just blocks writers


**Lock() and Unlock()** lock for writers. this waits for all readers to finish and blocks new ones from starting



Every lock unlock takes a few nanoseconds, by letting readers run in parrallel we can reclaim those nanoseconds back
## Step Goal: Refactor safestore to use RWMutex





# Step 3: Benchmarking 

Go has First class benchmarking features

ALl benchmarking tests start with Benchmark<funcname></funcname>


They take b *testing.B


We then run the test with go test -bench=specific test(optional)


Essential b methods

| Method   | What it does    |
|--------------- | --------------- |
| b.ResetTimer()   | Exclude setup time from measurement   |
| b.StopTimer()/ b.Starttimer()   | pause/resume clock mid loop   |
| Item1.3   | Item2.3   |
| Item1.4   | Item2.4   |

Sketch


naming convention is Benchmarkfunc


We can simulate multiple goroutines running inside
