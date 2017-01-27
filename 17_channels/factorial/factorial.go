package main

// example of parallel factorial computation using fan-in/fan-out pattern

import (
  "fmt"
)

func main() {
  in := gen()

  // fan-out
  c0 := factorial(in)
  c1 := factorial(in)
  c2 := factorial(in)
  c3 := factorial(in)
  c4 := factorial(in)
  c5 := factorial(in)
  c6 := factorial(in)
  c7 := factorial(in)
  c8 := factorial(in)
  c9 := factorial(in)

  // fan-in
  for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
    fmt.Println(n)
  }
}

func gen() <-chan int {
  out := make(chan int)
  go func() {
    for i := 0; i < 1000; i++ {
      for j := 3; j < 13; j++ {
        out <- j
      }
    }
    close(out)
  }()
  return out
}

func factorial(in <-chan int) <-chan int {
  out := make(chan int)
  go func() {
    for n := range in {
      out <- fact(n)
    }
    close(out)
  }()
  return out
}

func fact(n int) int {
  total := 1
  for i := n; i > 0; i-- {
    total *= i
  }
  return total
}

func merge(in ...<-chan int) <-chan int {
  done := make(chan bool)
  out := make(chan int)

  for _, c := range in {
    go func(d <-chan int) {
      for n := range d {
        out <- n
      }
      done <- true
    }(c)
  }

  go func() {
    for v := 0; v < len(in); v++ {
      <-done
    }
    close(done)
    close(out)
  }()

  return out
}
