package main

import "fmt"

// https://blog.golang.org/pipelines

func main() {
  // main() sets up the pipeline
  consumer(processing(generator()))
}

func generator() chan int {
  // the generator function, launches a go routine.
  out := make(chan int)
  go func() {
    for i := 0; i < 10; i++ {
      out <- i
    }
    close(out)
  }()
  return out
}

func processing(cx chan int) chan int {
  // the intermediate stages process the data launching a go routine.
  out := make(chan int)
  go func() {
    for v := range cx {
      out <- v * v
    }
    close(out)
  }()
  return out
}

func consumer(cx chan int) {
  // the final stage consumes the data
  for v := range cx {
    fmt.Println("Value:", v)
  }
}
