<!--markdownlint-disable-->


# Step 1: The Writer Interface and Mock


In GO, interfaces are satisfied implicitly, if a struct implements the methods it "is" that interface


## Step Goal: Define a LogWriter interface that accepts a batch of logs and create a MOckWriter struct that implements this interface. MockWriter should be able to track how many times it was called and what data it received


sketch:

type Log struct{
Level int
Message string
Timestamp time.Time
}


type LogWriter interface{
write([]Log Logs)
}


type MockWriter struct{
mu sync.Mutex
Logs []Log
}



MockWriter can store logs in a 


### Open/Closed Principle: Open to extension, closed to modificationof existing data
append(logs, batch...): ... is unpacking so we append each element individually instead of nesting slices in another slice


# Step 2: Naive Ingester



srv := http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request){})



# Step 3 : Monitor and Stress Test


## Sync/Atomic

Atomic Operations are lock free, thread-safe read writes on a single value. It is cheaper than a mutex when we just need to protect a single number


```Go

var counter int64

//update
atomic.AddInt64(&counter, 1)
atomic.AddInt64(&counter, -1)


//write
atomic.StoreInt64(&counter, 42)


//read
val := atomic.LoadInt64(&counter)




```


## Producer Consumer Model

We use a buffered channel so we can allow the handler to immedieately return even while waiting for slow database writes to happen


Producers(HTTP Handlers): Drop logs into a channel

Consumers(Workers): Pull  logs from the other end of the channel and do what they need to do with them


Standard Channel send operations will block the entire goroutine if the buffered channel is full. We use a select statement to make it non blockin


## Fixed Worker Pool


We start all our workers at once. Each worker runs a function to listen to logchan and proocesses

we can use for i := range i.logChan {}
to pull logs out of the channel one by one
