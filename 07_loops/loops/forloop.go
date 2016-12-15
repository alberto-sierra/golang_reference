package loops

import (
  "fmt"
  "time"
)

/*
The general form of a for loop is
for init; condition; post-action { }
*/

func Print10for() {
  for n := 1; n <= 10; n++ {
    fmt.Println(n)
  }
}

// an infinite for loop:
func Printforever(n int) {
  for {
    time.Sleep(5 * time.Millisecond)
    if n % 1000 == 0 {
      fmt.Printf("%d %s\n", n, "...press ctrl-c to stop...")
      time.Sleep(time.Second)
      n++
      continue // jumps to the next for iteration
    } else {
      fmt.Println(n)
      n++
    }
  }
}
