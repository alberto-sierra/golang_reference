package main

import (
  "fmt"
  "sync"
)

// "Do not communicate by sharing memory; instead, share memory by communicating."
func main() {
  // a channel is created with the make() helper function
  ch := make(chan int)
  cj := make(chan int)
  xt := make(chan bool) // a semaphore channel
  var wg sync.WaitGroup

  go func() {
    defer close(ch)
    for i := 0; i < 10; i++ {
      /*
         putting a value in a channel blocks execution until it is removed.
         To put a value IN, use the arrow pointing towards the channel.
      */
      ch <- i
    }
  }()

  /*
     fan-out pattern: single generator, multiple receivers
     using syng.WaitGroup to orchestrate processes
  */

  wg.Add(1)
  go func() {
    defer wg.Done()
    for v := range ch {
      // To take a value OUT, use the arrow with its back pointint to the channel.
      fmt.Println("foo:", v)
    }
  }()

  wg.Add(1)
  go func() {
    defer wg.Done()
    for v := range ch {
      fmt.Println("bar:", v)
    }
  }()

  /*
     fan-in pattern: multiple generator, single receiver.
     The channel must be closed after the generators are done.
     Using a semaphore channel to synchronize processes
  */
  go func() { // first generator: even numbers
    for i := 0; i < 10; i++ {
      if i%2 == 0 {
        cj <- i // blocks execution
      }
    }
    xt <- true // blocks execution
  }()

  go func() { // second generator: odd numbers
    for i := 0; i < 10; i++ {
      if i%2 != 0 {
        cj <- i // blocks execution
      }
    }
    xt <- true // blocks execution
  }()

  go func() {
    // this go function will terminate once the generators are done and close the channels, without blocking the execution of the receiver.
    <-xt // discards value from channel
    <-xt
    close(xt)
    close(cj)
  }()

  // receiver
  // range loops until the channel is closed (blocks execution)
  for v := range cj {
    fmt.Println("foobar:", v)
  }

  // wait for all go routines to finish
  wg.Wait()
}
