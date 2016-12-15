package main

import "fmt"

/*
Wrapper creates a local variable that is
remembered through runtime because it is stored
in a variable.
*/

func Wrapper() func() int {
  x := 0
  return func() int {
    x++
    return x
  }
}

/*
There is no function overloading in Go.
You cannot redefine a function with the same name,
even with different signature (name + params + return)
*/

func Wrapper2(x int) func() int {
  return func() int {
    x++
    return x
  }
}

func main() {

  // the variable increment "becomes" a function.
  var increment = Wrapper()
  fmt.Println(increment())
  fmt.Println(increment())

  x := 0
  increment = Wrapper2(x) // passing x as a value
  fmt.Println(increment())
  fmt.Println(increment())

}
