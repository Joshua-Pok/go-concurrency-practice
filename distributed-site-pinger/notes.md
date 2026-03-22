<!--markdownlint-disable-->


# Step 1: Atomic Pinger

net/http is standard tool for making http requests
for pinging a site, a HEAD request is generally preferred over GET because it only retrieves headers, saving time and bandwidht

## Step Goal: Create a function that takes a URL and returns a summary of results, including status code, time taken and any errors


Sketch:


type Result struct{
statusCode:
Headers:
timeTaken:
Errors:
}

func pingSite(URL string) Result{

start := time.Now()
res, err := http.Head(URL)
if err != nil{
return err
}

defer res.Body.Close()

elapsed := time.Since(start)

serialize res into Result struct
}


# Step 2: The Bounded Coordinator


## Sephamore Pattern
Semaphore is a synchronization pattenr used t omanage access to shared resources using a counter.


When threads/processes want to access shared resources, they "acquire" semaphore. If counter is 0 they wait until another thread releases it, regulating access

In Go, a **bufferred channel is the idiomatic way of building a spehamore.



## Step Goal: Implement a Monitor function that takes a lsice of URLs and a concurrency limit, It should return a channel where results are streamed back as they finish


sketch:

func Monitor(urls strings, limit int) <- chan Result{

wg := sync.WaitGroup{}
sem := make(chan struct{}, limit) // channel of channels

resultsChan := make(chan Result, len(urls))

//we use a resultsChan so urls respond time dont block 

//reason why we use a empty struct is because its 0 bytes, bool is one bye


for i, url := range(urls){
wg.Add(1)

sem <- struct{}{} //we try to "acquire" access, will block if the channel is full
go func (i int, url string){
defer  wg.Done()
defer func(){<-sem}() release slot
resultsChan <- pingSite(url) 
}(i, url)

}


go func(){
wg.Wait()
close(resultsChan)
}

return resultsChan
}



**LIfecycle of goroutines ARE NOT TIED to fucntion that spawned them. They live until they return**


## Streaming Vs Collection

It is better to return a channel instead of collecting into a list so that we dont have to wait for the slowest routine tofihish before seeing any data


# Code Review

🔍 Summary
──────────────
  • Empty main.go makes application non-functional
  • Critical bug in Monitor function's goroutine management
  • Go version 1.26.1 doesn't exist
  • Error handling incomplete in Monitor function

⚠️  Issues
─────────
  1. ● [HIGH] main.go contains only package declaration with no actual implementation — distributed-site-pinger/cmd/pinger/main.go:1
  2. ● [HIGH] Monitor function incorrectly launches a goroutine per URL that waits for all URLs to complete and closes results channel prematurely — distributed-site-pinger/internal/monitor/pinger.go:34
  3. ● [HIGH] Go version 1.26.1 doesn't exist; should be 1.21.x or similar — distributed-site-pinger/go.mod:3
  4. ● [HIGH] Errors from pingSite are ignored with _, causing silent failures — distributed-site-pinger/internal/monitor/pinger.go:32

🧹 Code Quality
───────────────
  1. ◐ [MEDIUM] Monitor function overly complex; simplify by removing redundant goroutine and fixing waitgroup usage — distributed-site-pinger/internal/monitor/pinger.go
  2. ○ [LOW] Variable names could be more descriptive (e.g., 'sem' instead of 'semaphore' or 'concurrencyLimit') — distributed-site-pinger/internal/monitor/pinger.go

🚀 Performance
──────────────
  1. ◐ [MEDIUM] Results channel created with capacity equal to URL count, which could be memory-intensive for large lists — distributed-site-pinger/internal/monitor/pinger.go:18

🔐 Security
───────────
  1. ◐ [MEDIUM] No URL validation before pinging could lead to security issues with untrusted input — distributed-site-pinger/internal/monitor/pinger.go

🧪 Testing Suggestions
──────────────────────
  1. ◐ [MEDIUM] TestSmoke is a placeholder that always passes and provides no value — distributed-site-pinger/internal/monitor/pinger_test.go:5
  2. ◐ [MEDIUM] TestMonitorConcurrency doesn't actually test any concurrency behavior — distributed-site-pinger/internal/monitor/pinger_test.go:30
  3. ○ [LOW] No tests for error cases or timeout scenarios — distributed-site-pinger/internal/monitor/pinger_test.go

✅ Positives
────────────
  1. • pingSite function properly uses http.Head, tracks timing, and closes response body
  2. • Correctly uses semaphore pattern with buffered channel for concurrency control
  3. • Uses channels for streaming results instead of collecting all before returning

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  ⚠  Review complete — 4 high · 6 total issues
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━






# Step 3: Naive Timeout


Server might not fail immediately. It might just sit there and do nothing. If a pinger keeps waiting for a dead url. it wastes a slot in our semaphore.


go select statement allows a goroutine to wait on multiple communication operations


select statements must be **channel** operations

time.After duration returns a channel that sends a value after specified time has elapsed


by racing our pinger result channel against time.after we can implement a timeout

## Step Goal: Update Worker Logic so that if pingSite takes longer than a specific duration. worker returns a timeout error instead of waiting indefinitely


Sketch

define MAX_WAIT_TIME constant = 30

go func(i int, url string){

defer wg.Done()
defer func() {<- sem }()


start = time.Now()


timeout := time.After(true)


select{

case timeout:
<-sem

case pingSite(url):
if err != nil{
result.Error = err

}


results <- result
}


result, err := pingSite(url)
}

