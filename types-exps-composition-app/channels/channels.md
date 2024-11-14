# Channels

***Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.***


Create a new channel with _make(chan val-type)_. Channels are typed by the values they convey.
```
messages := make(chan string)
```

_Send_ a value into a channel using the channel **_<-_** syntax. Send _`"ping"`_ to the messages channel we made above, from a new goroutine.
```
go func() { messages <- "ping" }()
```

The `_<-channel_` syntax receives a value from the channel. Here we’ll receive the _`"ping"`_ message we sent above and print it out.
```
msg := <-messages
fmt.Println(msg)
```

When the program is executed, the _`"ping"`_ message is successfully passed from one goroutine to another via channel.
```
$ go run channels.go 
ping
```

**NOTE:** By default sends and receives block until both the sender and receiver are ready. This property allows the user to wait at the end of the program for the _`"ping"`_ message without having to use any other synchronization.

## Channel Buffering

By default channels are _unbuffered_, meaning that they will only accept sends `(chan <-)` if there is a corresponding receive `(<- chan)` ready to receive the sent value. _Buffered channels_ accept a limited number of values without a corresponding receiver for those values.

_`make`_ a channel of strings buffering up to 2 values
```
messages := make(chan string, 2)
```

Because this channel is buffered, send these values into the channel without a corresponding concurrent receive.
```
messages <- "buffered"
messages <- "channel"
```

Later, receive these two values as usual.
```
fmt.Println(<-messages)
fmt.Println(<-messages)
```

Post execution of the program:
```
$ go run channel-buffering.go 
buffered
channel
```

## Channel Synchronization

Use channels to synchronize execution across goroutines. Here's an example of using a blocking receivie to wait for a goroutine to finish. When waiting for multiple goroutines to finish, prefer to use a <ins>WaitGroup</ins>

In this function the `done` channel will be used to notify another goroutine that this function's work is done. Send a value to notify that this function is done - `done <- true`
```
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    done <- true
}
```

In the main function, start a worker goroutine, giving it the channel to notify on.
```
done := make(chan bool, 1)
go worker(done)
```

Block until a notification is received from the worker on the channel in main function.
```
<-done
```

Post executing the program
```
$ go run channel-synchronization.go      
working...done 
```

**NOTE:** If the `<-done` is removed from the program, then the program would exit before the worker even started.

## Channel Directions

When using channels as function parameters, user can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.

This `ping` function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
```
func ping(pings chan<- string, msg string) {
    pings <- msg
}
```

The pong function accepts one channel to receive `(pings)` and a second to send `(pongs)`.
```
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
```

The main function would be used to call both the functions
```
func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
```

## Select

Go’s _select_ let's the user wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

In the below example let's select across two channels. Add below code in the main function.
```
c1 := make(chan string)
c2 := make(chan string)
```

Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
```
go func() {
    time.Sleep(1 * time.Second)
    c1 <- "one"
}()
go func() {
    time.Sleep(2 * time.Second)
    c2 <- "two"
}()
```

Let's use _select_ to await both of these values simultaneously, printing each one as it arrives.
```
for i := 0; i < 2; i++ {
    select {
    case msg1 := <-c1:
        fmt.Println("received", msg1)
    case msg2 := <-c2:
        fmt.Println("received", msg2)
    }
}
```

The values received would be `"one"` and then `"two"` as expected.
```
$ time go run select.go 
received one
received two
```

Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.
```
real    0m2.245s
```

## Timeouts

_Timeouts_ are important for programs that connect to external resources or that otherwise need to bound execution time. Implementing timeouts in Go is easy and elegant thanks to channels and `select`.

Add below code in the main function.

In this example, suppose user executes an external call that returns its result on a channel c1 after 2s. Note that the channel is buffered, so the send in the goroutine is nonblocking. This is a common pattern to prevent goroutine leaks in case the channel is never read.
```
c1 := make(chan string, 1)
go func() {
    time.Sleep(2 * time.Second)
    c1 <- "result 1"
}()
```

Here’s the `select` implementing a timeout. `res := <-c1` awaits the result and `<-time.After` awaits a value to be sent after the timeout of 1s. Since select proceeds with the first receive that’s ready, let's take the timeout case if the operation takes more than the allowed 1s.
```
select {
case res := <-c1:
    fmt.Println(res)
case <-time.After(1 * time.Second):
    fmt.Println("timeout 1")
}
```

If allowed a longer timeout of 3s, then the receive from c2 will succeed and result will be printed.
```
c2 := make(chan string, 1)
go func() {
    time.Sleep(2 * time.Second)
    c2 <- "result 2"
}()
select {
case res := <-c2:
    fmt.Println(res)
case <-time.After(3 * time.Second):
    fmt.Println("timeout 2")
}
```

Running this program shows the first operation timing out and the second succeeding.
```
$ go run timeouts.go 
timeout 1
result 2
```

## Non-Blocking Channel Operations

Basic sends and receives on channels are blocking. However, user can use `select` with a `default` clause to implement _non-blocking_ sends, receives, and even non-blocking multi-way `select` cases.

Here’s a non-blocking receive. If a value is available on `messages` then `select` will take the `<-messages case` with that value. If not it will immediately take the `default` case.
```
select {
case msg := <-messages:
    fmt.Println("received message", msg)
default:
    fmt.Println("no message received")
}
```

A non-blocking send works similarly. Here msg cannot be sent to the `messages` channel, because the channel has no buffer and there is no receiver. Therefore the `default` case is selected.
```
msg := "hi"
select {
case messages <- msg:
    fmt.Println("sent message", msg)
default:
    fmt.Println("no message sent")
}
```

User can use multiple `cases` above the `default` clause to implement a multi-way non-blocking select. In this example, let's attempt non-blocking receives on both `messages` and `signals`.
```
select {
case msg := <-messages:
    fmt.Println("received message", msg)
case sig := <-signals:
    fmt.Println("received signal", sig)
default:
    fmt.Println("no activity")
}
```

Post executing the program
```
$ go run non-blocking-channel-operations.go 
no message received
no message sent
no activity
```


## Closing Channels


_Closing_ a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.

In this example let's use a `jobs` channel to communicate work to be done from the `main()` goroutine to a worker goroutine. When there are no more jobs for the worker then the `jobs` channel is `closed`.
```
jobs := make(chan int, 5)
done := make(chan bool)
```

The worker goroutine repeatedly receives from `jobs` with j, more `:= <-jobs`. In this special 2-value form of receive, the `more` value will be `false` if `jobs` has been `closed` and all values in the channel have already been received. Let's use this to notify on `done` when all the jobs are completed.
```
go func() {
    for {
        j, more := <-jobs
        if more {
            fmt.Println("received job", j)
        } else {
            fmt.Println("received all jobs")
            done <- true
            return
        }
    }
}()
```

This sends 3 jobs to the worker over the `jobs` channel, then closes it.
```
for j := 1; j <= 3; j++ {
    jobs <- j
    fmt.Println("sent job", j)
}
close(jobs)
fmt.Println("sent all jobs")
```

Await the worker using the <ins>[synchronization](https://github.com/joshi-ani/learn_go/blob/main/types-exps-composition-app/channels/channels.md#channel-synchronization)</ins> approach used earlier.
```
<-done
```

Reading from a closed channel succeeds immediately, returning the zero value of the underlying type. The optional second return value is true if the value received was delivered by a successful send operation to the channel, or `false` if it was a zero value generated because the channel is closed and empty.
```
_, ok := <-jobs
fmt.Println("received more jobs:", ok)
```

## Range over Channels

Let's iterate over 2 values in the queue channel.
```
queue := make(chan string, 2)
queue <- "one"
queue <- "two"
close(queue)
```

This `range` iterates over each element as it’s received from `queue`. Because the channel used above is `closed` and the iteration terminates after receiving the 2 elements.
```
for elem := range queue {
    fmt.Println(elem)
}
```
```
$ go run range-over-channels.go
one
two
```

This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.

## Timers

Timers represent a single event in the future. User can set the timer for how long it can wait, and it provides a channel that will be notified at that time. This timer will wait 2 seconds.
```
timer1 := time.NewTimer(2 * time.Second)
```

The `<-timer1.C` blocks on the timer’s channel `C` until it sends a value indicating that the timer fired.
```
<-timer1.C
fmt.Println("Timer 1 fired")
```

One reason a timer may be useful is that user can cancel the timer before it fires. Here’s an example of that.
```
timer2 := time.NewTimer(time.Second)
go func() {
    <-timer2.C
    fmt.Println("Timer 2 fired")
}()
stop2 := timer2.Stop()
if stop2 {
    fmt.Println("Timer 2 stopped")
}
```

Give the `timer2` enough time to fire, if it ever was going to, to show it is in fact stopped.
```
time.Sleep(2 * time.Second)
```


The first timer will fire ~2s after we start the program, but the second should be stopped before it has a chance to fire.
```
$ go run timers.go
Timer 1 fired
Timer 2 stopped
```

## Tickers


<ins>Timers</ins> are for when user wants to do something once in the future - tickers are for when user wants to do something repeatedly at regular intervals. Here’s an example of a ticker that ticks periodically until it's stopped.

Tickers use a similar mechanism to timers: a channel that is sent values. 
In this example let's use the `select` builtin on the channel to await the values as they arrive every 500ms

```
ticker := time.NewTicker(500 * time.Millisecond)
done := make(chan bool)

go func() {
    for {
        select {
        case <-done:
            return
        case t := <-ticker.C:
            fmt.Println("Tick at", t)
        }
    }
}()
```

Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel. In the above example let's stop the ticker after 1600ms.
```
time.Sleep(1600 * time.Millisecond)
ticker.Stop()
done <- true
fmt.Println("Ticker stopped")
```

Post running the program, the ticker should tick 3 times before it's stopped.
```
$ go run tickers.go
Tick at 2012-09-23 11:29:56.487625 -0700 PDT
Tick at 2012-09-23 11:29:56.988063 -0700 PDT
Tick at 2012-09-23 11:29:57.488076 -0700 PDT
Ticker stopped
```

## Worker Pools


In this example let's look at how to implement a worker pool using goroutines and channels.

The worker will run several concurrent instances. These workers will receive work on the jobs channel and send the corresponding results on results. Let's sleep a second per job to simulate an expensive task.
```
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}
```

In order to use the pool of workers, lets make 2 channels to send them work and collect the results.
```
for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
}
```

Send 5 `jobs` and then `close` that channel to indicate that’s all the work we have.
```
for j := 1; j <= numJobs; j++ {
    jobs <- j
}
close(jobs)
```

Finally, collect all the results of the work. This also ensures that the worker goroutines have finished. An alternative way to wait for multiple goroutines is to use a <ins>[WaitGroup](https://github.com/joshi-ani/learn_go/blob/main/types-exps-composition-app/channels/channels.md#wait-groups)</ins>.
```
for a := 1; a <= numJobs; a++ {
    <-results
}
```

Post running the program shows 5 jobs being executed by various workers. The program only takes about 2 seconds despite doing about 5 seconds of total work becuase there are 3 workers operating concurrently.
```
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5
real    0m2.358s
```

## Wait Groups

To wait for multiple goroutines to finish, a _wait group_ is used.

Following function will be executed in every goroutine. Sleep is used to simulate an expensive task.
```
func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}
```

WaitGroup is used to wait for all the goroutines launched here to finish. Note: if a WaitGroup is explicitly passed into functions, it should be done by _pointer_.
```
var wg sync.WaitGroup
```

Launch several goroutines and increment the WaitGroup counter for each. Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done. This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
```
for i := 1; i <= 5; i++ {
    wg.Add(1)

    go func() {
        defer wg.Done()
        worker(i)
    }()
}
```

Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
```
wg.Wait()
```


Note that this approach has no straightforward way to propagate errors from workers. For more advanced use cases, consider using the <ins>[errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup)</ins>.

The order of workers starting up and finishing is likely to be different for each invocation.
```
$ go run waitgroups.go
Worker 5 starting
Worker 3 starting
Worker 4 starting
Worker 1 starting
Worker 2 starting
Worker 4 done
Worker 1 done
Worker 2 done
Worker 5 done
Worker 3 done
```

## Rate Limiting


<ins>[Rate limiting](https://en.wikipedia.org/wiki/Rate_limiting)</ins> is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and <ins>[tickers](https://github.com/joshi-ani/learn_go/blob/main/types-exps-composition-app/channels/channels.md#tickers)</ins>.

Suppose a limit needs to be added for handling incoming requests. Let's server these requests off a channel for the same name.
```
requests := make(chan int, 5)
for i := 1; i <= 5; i++ {
    requests <- i
}
close(requests)
```

This `limiter` channel will receive a value every 200 milliseconds. This is the regulator in the rate limiting scheme defined above.
```
limiter := time.Tick(200 * time.Millisecond)
```

By blocking on a receive from the limiter channel before serving each request, let's limit this to 1 request every 200 milliseconds.
```
for req := range requests {
    <-limiter
    fmt.Println("request", req, time.Now())
}
```

Let's allow short bursts of requests in this rate limiting scheme while preserving the overall rate limit. This can be accomplished by buffering the limiter channel. This burstyLimiter channel will allow bursts of up to 3 events.
```
burstyLimiter := make(chan time.Time, 3)
```

Fill up the channel to represent allowed bursting.
```
for i := 0; i < 3; i++ {
    burstyLimiter <- time.Now()
}
```

Every 200 milliseconds let's try to add a new value to burstyLimiter, up to its limit of 3.
```
go func() {
    for t := range time.Tick(200 * time.Millisecond) {
        burstyLimiter <- t
    }
}()
```

Now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability of burstyLimiter.
```
burstyRequests := make(chan int, 5)
for i := 1; i <= 5; i++ {
    burstyRequests <- i
}
close(burstyRequests)
for req := range burstyRequests {
    <-burstyLimiter
    fmt.Println("request", req, time.Now())
}
```

Post running the program, the first batch of requests handled once every ~200 milliseconds as desired.
```
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC
```
For the second batch of requests, the first 3 are immediately served because of the burstable rate limiting, then the remaining 2 are served with ~200ms delays each.
```
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
```

## Atomic Counters

The primary mechanism for managing state in Go is communication over channels. There are a few other options for managing state though. Lets look at using the `sync/atomic` package for atomic counters accessed by multiple goroutines.

Let's use an atomic integer type to represent our (always positive) counter.
```
package main
import (
    "fmt"
    "sync"
    "sync/atomic"
)
func main() {
    var ops atomic.Uint64
}
```
A WaitGroup will help to wait for all goroutines to finish their work.
```
var wg sync.WaitGroup
```

Let's start 50 goroutines that each increment the counter exactly 1000 times. To atomically increment the counter, use `Add`. 
```
for i := 0; i < 50; i++ {
    wg.Add(1)
    go func() {
        for c := 0; c < 1000; c++ {
            ops.Add(1)
        }
        wg.Done()
    }()
}
```

Wait until all the goroutines are done.
```
wg.Wait()
```

Here no goroutines are writing to ‘ops’, but using `Load` it’s safe to atomically read a value even while other goroutines are (atomically) updating it.
```
fmt.Println("ops:", ops.Load())
```

Expect to get exactly 50,000 operations. If a non-atomic integer had been used and incremented it with ops++, we’d likely get a different number, changing between runs, because the goroutines would interfere with each other. Moreover, we’d get data race failures when running with the -race flag.
```
$ go run atomic-counters.go
ops: 50000
```


## Mutexes

For more complex state a _mutex_ can be used to safely access data across multiple goroutines.

Container holds a map of counters; since we want to update it concurrently from multiple goroutines, let's add a `Mutex` to synchronize access. **Note** that mutexes must not be copied, so if this _struct_ is passed around, it should be done by pointer.
```
type Container struct {
    mu       sync.Mutex
    counters map[string]int
}
```
Lock the mutex before accessing `counters`; unlock it at the end of the function using a `defer` statement.
```
func (c *Container) inc(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}
```

Note that the zero value of a mutex is usable as-is, so no initialization is required here. The doIncrement function increments a named counter in a loop.
```
func main() {
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }
    var wg sync.WaitGroup
    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }
}
```
Run several goroutines concurrently; note that they all access the same `Container`, and two of them access the same counter. Wait for the goroutines to finish.
```
wg.Add(3)
go doIncrement("a", 10000)
go doIncrement("a", 10000)
go doIncrement("b", 10000)
wg.Wait()
fmt.Println(c.counters)
```


Running the program shows that the counters updated as expected.
```
$ go run mutexes.go
map[a:20000 b:10000]
```

## Stateful Goroutines


In the previous example we used explicit locking with <ins>[mutexes](https://github.com/joshi-ani/learn_go/blob/main/types-exps-composition-app/channels/channels.md#mutexes)</ins> to synchronize access to shared state across multiple goroutines. Another option is to use the built-in synchronization features of goroutines and channels to achieve the same result. This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.

In this example our state will be owned by a single goroutine. This will guarantee that the data is never corrupted with concurrent access. In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies. These readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to respond.
```
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}
```

Let's count how many operations are performed.

The `reads` and `writes` channels will be used by other goroutines to issue read and write requests, respectively.
```
reads := make(chan readOp)
writes := make(chan writeOp)
```

Here is the goroutine that owns the `state`, which is a map as in the previous example but now private to the stateful goroutine. This goroutine repeatedly selects on the `reads` and `writes` channels, responding to requests as they arrive. A response is executed by first performing the requested operation and then sending a value on the response channel `resp` to indicate success (and the desired value in the case of `reads`).
```
go func() {
    var state = make(map[int]int)
    for {
        select {
        case read := <-reads:
            read.resp <- state[read.key]
        case write := <-writes:
            state[write.key] = write.val
            write.resp <- true
        }
    }
}()
```

This starts 100 goroutines to issue reads to the state-owning goroutine via the `reads` channel. Each read requires constructing a `readOp`, sending it over the `reads` channel, and then receiving the result over the provided `resp` channel.

```
for r := 0; r < 100; r++ {
    go func() {
        for {
            read := readOp{
                key:  rand.Intn(5),
                resp: make(chan int)}
            reads <- read
            <-read.resp
            atomic.AddUint64(&readOps, 1)
            time.Sleep(time.Millisecond)
        }
    }()
}
```

Let's start 10 writes as well, using a similar approach.
```
for w := 0; w < 10; w++ {
    go func() {
        for {
            write := writeOp{
                key:  rand.Intn(5),
                val:  rand.Intn(100),
                resp: make(chan bool)}
            writes <- write
            <-write.resp
            atomic.AddUint64(&writeOps, 1)
            time.Sleep(time.Millisecond)
        }
    }()
}
```

Let the goroutines work for a second.
```
time.Sleep(time.Second)
```

Finally, capture and report the op counts.
```
readOpsFinal := atomic.LoadUint64(&readOps)
fmt.Println("readOps:", readOpsFinal)
writeOpsFinal := atomic.LoadUint64(&writeOps)
fmt.Println("writeOps:", writeOpsFinal)
```

Running our program shows that the goroutine-based state management example completes about 80,000 total operations.
```
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177
```

For this particular case the goroutine-based approach was a bit more involved than the mutex-based one. It might be useful in certain cases though, for example when other channels are involved or when managing multiple such mutexes would be error-prone. Use whichever approach feels most natural, especially with respect to understanding the correctness of the program.
