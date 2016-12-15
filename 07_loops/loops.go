package main

import (
  "fmt"
  "github.com/alberto-sierra/golang_reference/07_loops/loops"
  "time"
  )

func main() {
  fmt.Println("A for loop:")
  loops.Print10for()
  fmt.Println("A for loop that mimics a while loop:")
  loops.Print10while()
  fmt.Println("A for loop that mimics a do-while loop:")
  stopat := 5
  loops.Printuntil(stopat)
  fmt.Println("An (almost) infinite for loop:")
  time.Sleep(3 * time.Second)
  startat := 1
  loops.Printforever(startat)
}
