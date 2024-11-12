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

# Channel Buffering

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

# Channel Synchronization

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

# Channel Directions

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