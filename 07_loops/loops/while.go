package loops

import "fmt"

/*
There is no do-while in go.
The while loop is achieved also with a for:
for condition { }
*/

func Print10while() {
  n := 1
  for n <= 10 {
    fmt.Println(n)
    n++
  }
}

func Printuntil(stopat int) {
  n := 1
  for {
    if n <= stopat {
      fmt.Println(n)
      n++
    } else {
      break // terminates the loop
    }

  }

}
