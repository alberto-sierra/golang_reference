package main

import (
  "fmt"
  "sync"
  "sync/atomic"
)

/*
  A WaitGroup waits for a collection of goroutines to finish.
  The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished.
*/
var wg, wg1, wg2, wg3, wg4 sync.WaitGroup
var c1, c2 int
var a1 int64
var mutex sync.Mutex

func main()  {
  /*
    Concurrency is managed with Go routines
    Because main() is also a go routine, the following threads
    won't execute because main() won't wait for them to finish,
    unless something makes the main() go routine wait (a waitgroup or a channel)
  */
  go one()
  go two()
  /*
  "Add" adds delta to the WaitGroup counter. If the counter becomes zero, all goroutines blocked on Wait are released.
  */
  wg.Add(2)
  go three()
  go four()
  // "Wait" blocks until the WaitGroup counter is zero. Wait can be used to block until all goroutines have finished.
  wg.Wait()

  // Race conditions can occur when using a shared resource with different threads: the total count here should be 20:
  wg.Add(20)
  for i := 0; i < 20; i++ {
    go func() {
      defer wg.Done()
      counter()
    }()
  }
  fmt.Println("Final count (race conditions): ", c1)
  wg.Wait()


  // The same using a mutex: a mutual exclusion lock
  wg.Add(20)
  for j := 0; j < 20; j++ {
    go func() {
      defer wg.Done()
      counterMutex()
    }()
  }
  wg.Wait()
  fmt.Println("Final count (with mutex): ", c2)

  // Again, using atomicity (see counterAtomic())
  wg.Add(20)
  for k := 0; k < 20; k++ {
    go func() {
      counterAtomic()
      wg.Done()
    }()
  }
  wg.Wait()
  fmt.Println("Final count (with atomicity): ", a1)
}

func one() {
  fmt.Println("one")
}

func two() {
  fmt.Println("two")
}

func three() {
  // Done decrements the WaitGroup counter.
  defer wg.Done()
  fmt.Println("three")
}

func four() {
  defer wg.Done()
  fmt.Println("four")
}

func counter() {
  c1++
}

func counterMutex() {
  //  "Lock" locks c2. If the lock is already in use, the calling goroutine blocks until the mutex is available.
  mutex.Lock()
  c2++
  // "Unlock"  unlocks c2. It is a run-time error if c2 is not locked on entry to Unlock.
  mutex.Unlock()
}

func counterAtomic() {
   // Package atomic provides low-level atomic memory primitives.
   // AddInt64 atomically adds delta to *addr and returns the new value.
  atomic.AddInt64(&a1, 1)
}
